package models

import (
	"time"

	"github.com/zmb3/spotify/v2"
)

type Playlist struct {
	ID        int64      `json:"artist_id,omitempty"`
	SpotifyID spotify.ID `json:"spotify_id"`
	Title     string     `json:"Title"`
	Monitor   bool       `json:"monitor"`
	DateAdded time.Time  `json:"date_added,omitempty"`
	Tracks    []Track    `json:"tracks"`
}
