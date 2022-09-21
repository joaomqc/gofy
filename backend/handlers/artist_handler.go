package handlers

import (
	"gofy/forms"
	"gofy/models"
	spotifyInfra "gofy/spotify"
)

type ArtistHandler struct{}

var albumHandler = new(AlbumHandler)

func (h ArtistHandler) Add(form forms.AddArtist) (*models.Artist, error) {
	client, ctx := spotifyInfra.GetClient()
	result, err := client.GetArtist(ctx, form.SpotifyId)

	if err != nil {
		return nil, err
	}

	artist := &models.Artist{
		SpotifyID: result.ID,
		Name:      result.Name,
		Monitor:   form.Monitor,
	}

	err = artist.Store()

	if err != nil {
		return nil, err
	}

	return artist, nil
}
