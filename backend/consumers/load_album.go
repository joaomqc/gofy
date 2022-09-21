package consumers

import (
	"errors"
	"gofy/db"
	"gofy/handlers"
	"gofy/messages"
	"gofy/models"
)

type LoadAlbumConsumer struct{}

var trackHandler = new(handlers.TrackHandler)

func (c LoadAlbumConsumer) HandleLoadAlbumMessage(message messages.LoadAlbum) error {
	sqlStatement := `
		SELECT id, spotify_id, monitor
		FROM album
		WHERE album.id = ?
	`

	row := db.GetDB().QueryRow(sqlStatement, message.AlbumId)

	if row == nil {
		return errors.New("album not found")
	}

	album := models.Album{}
	err := row.Scan(&album.ID, &album.SpotifyID, &album.Monitor)

	if err != nil {
		return err
	}

	err = trackHandler.LoadFromAlbum(album)

	return err
}
