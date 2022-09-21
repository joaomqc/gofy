package models

import (
	"gofy/db"
	"gofy/infrastructure"
	"gofy/messages"
	"time"

	"github.com/zmb3/spotify/v2"
)

type Track struct {
	ID        int64      `json:"track_id,omitempty"`
	SpotifyID spotify.ID `json:"spotify_id"`
	AlbumID   int64      `json:"album_id"`
	Title     string     `json:"title"`
	Monitor   bool       `json:"monitor"`
	DateAdded time.Time  `json:"date_added,omitempty"`
}

func (t *Track) Store() error {
	id, err := t.insert()
	if err != nil {
		return err
	}

	t.ID = *id

	if t.Monitor {
		err = t.publishMessages()
		if err != nil {
			return err
		}
	}

	return nil
}

func (t Track) insert() (*int64, error) {
	sqlStatement := `
		INSERT INTO track (spotify_id, title, album_id, monitor)
		VALUES ($1, $2, $3, $4)
	`
	result, err := db.GetDB().Exec(sqlStatement, t.SpotifyID, t.Title, t.AlbumID, t.Monitor)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (t Track) publishMessages() error {
	return infrastructure.Publish(messages.DownloadTrack{
		TrackId: t.ID,
	}, infrastructure.DownloadTrackQueue)
}
