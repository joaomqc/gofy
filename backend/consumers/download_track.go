package consumers

import (
	"errors"
	"gofy/db"
	"gofy/messages"
	"gofy/models"
)

type DownloadTrackConsumer struct{}

func (c DownloadTrackConsumer) HandleDownloadTrackMessage(message messages.DownloadTrack) error {
	sqlStatement := `
		SELECT id, spotify_id, monitor
		FROM track
		WHERE track.id = ?
	`

	row := db.GetDB().QueryRow(sqlStatement, message.TrackId)

	if row == nil {
		return errors.New("track not found")
	}

	track := models.Track{}
	err := row.Scan(&track.ID, &track.SpotifyID, &track.Monitor)

	if err != nil {
		return err
	}

	err = trackHandler.DownloadTrack(track)

	return err
}
