package forms

import "github.com/zmb3/spotify/v2"

type AddArtist struct {
	SpotifyId spotify.ID `json:"spotify_id"`
	Monitor   bool       `json:"monitor"`
}
