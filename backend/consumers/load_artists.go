package consumers

import (
	"errors"
	"gofy/db"
	"gofy/handlers"
	"gofy/messages"
	"gofy/models"
)

type LoadArtistConsumer struct{}

var albumHandler = new(handlers.AlbumHandler)

func (c LoadArtistConsumer) HandleLoadArtistMessage(message messages.LoadArtist) error {
	sqlStatement := `
		SELECT id, spotify_id, monitor
		FROM artist
		WHERE artist.id = ?
	`

	row := db.GetDB().QueryRow(sqlStatement, message.ArtistId)

	if row == nil {
		return errors.New("artist not found")
	}

	artist := models.Artist{}
	err := row.Scan(&artist.ID, &artist.SpotifyID, &artist.Monitor)

	if err != nil {
		return err
	}

	err = albumHandler.LoadFromArtist(artist)

	return err
}
