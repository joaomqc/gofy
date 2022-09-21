package models

import (
	"gofy/db"
	"gofy/infrastructure"
	"gofy/messages"
	"time"

	"github.com/zmb3/spotify/v2"
)

type Album struct {
	ID          int64      `json:"album_id,omitempty"`
	SpotifyID   spotify.ID `json:"spotify_id"`
	ArtistId    int64      `json:"artist_id"`
	Title       string     `json:"title"`
	AlbumType   string     `json:"album_type"`
	Monitor     bool       `json:"monitor"`
	ReleaseDate time.Time  `json:"release_date"`
	DateAdded   time.Time  `json:"date_added,omitempty"`
	Tracks      []Track    `json:"tracks"`
}

func (a *Album) Store() error {
	id, err := a.insert()
	if err != nil {
		return err
	}

	a.ID = *id

	if a.Monitor {
		err = a.publishMessages()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a Album) insert() (*int64, error) {
	sqlStatement := `
		INSERT INTO album (spotify_id, title, artist_id, album_type, monitor, release_date)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	result, err := db.GetDB().Exec(sqlStatement, a.SpotifyID, a.Title, a.ArtistId, a.AlbumType, a.Monitor, a.ReleaseDate)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (a Album) publishMessages() error {
	return infrastructure.Publish(messages.LoadAlbum{
		AlbumId: a.ID,
	}, infrastructure.LoadAlbumQueue)
}
