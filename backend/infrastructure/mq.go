package infrastructure

import (
	"encoding/json"

	"github.com/adjust/rmq/v4"
)

var LoadArtistQueue rmq.Queue
var LoadAlbumQueue rmq.Queue
var DownloadTrackQueue rmq.Queue

func Publish(message interface{}, queue rmq.Queue) error {
	bytes, err := json.Marshal(message)

	if err != nil {
		return err
	}

	err = queue.PublishBytes(bytes)
	if err != nil {
		return err
	}

	return nil
}

func InitQueues(connection rmq.Connection) error {
	var err error
	LoadArtistQueue, err = connection.OpenQueue("load_artist")
	if err != nil {
		return err
	}

	LoadAlbumQueue, err = connection.OpenQueue("load_album")
	if err != nil {
		return err
	}

	DownloadTrackQueue, err = connection.OpenQueue("download_track")

	return err
}
