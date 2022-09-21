package models

import (
	"gofy/db"
	"gofy/infrastructure"
	"gofy/messages"
	"time"

	"github.com/zmb3/spotify/v2"
)

type Artist struct {
	ID        int64      `json:"artist_id,omitempty"`
	SpotifyID spotify.ID `json:"spotify_id"`
	Name      string     `json:"name"`
	Monitor   bool       `json:"monitor"`
	DateAdded time.Time  `json:"date_added,omitempty"`
	Albums    []Album    `json:"albums,omitempty"`
}

func (a *Artist) Store() error {
	id, err := a.insert()

	if err != nil {
		return err
	}

	a.ID = *id

	err = a.publishMessages()

	return err
}

func (a Artist) insert() (*int64, error) {
	sqlStatement := `
		INSERT INTO artist (spotify_id, name, monitor)
		VALUES ($1, $2, $3)
	`
	result, err := db.GetDB().Exec(sqlStatement, a.SpotifyID, a.Name, a.Monitor)

	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (a Artist) publishMessages() error {
	return infrastructure.Publish(messages.LoadArtist{
		ArtistId: a.ID,
	}, infrastructure.LoadArtistQueue)
}
