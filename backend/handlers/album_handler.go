package handlers

import (
	"gofy/models"
	spotifyInfra "gofy/spotify"

	"github.com/zmb3/spotify/v2"
)

type AlbumHandler struct{}

var albumTypes = []spotify.AlbumType{
	spotify.AlbumTypeAlbum,
	spotify.AlbumTypeSingle,
}

func (h AlbumHandler) LoadFromArtist(artist models.Artist) error {
	client, ctx := spotifyInfra.GetClient()
	albumsPage, err := client.GetArtistAlbums(ctx, artist.SpotifyID, albumTypes, spotify.Market("PT"))

	if err != nil {
		return err
	}

	spotifyAlbums := albumsPage.Albums

	for {
		if albumsPage.Next == "" {
			break
		}

		err = client.NextPage(ctx, albumsPage)
		if err != nil {
			return err
		}
		spotifyAlbums = append(spotifyAlbums, albumsPage.Albums...)
	}

	for _, spotifyAlbum := range spotifyAlbums {
		album := models.Album{
			SpotifyID:   spotifyAlbum.ID,
			ArtistId:    artist.ID,
			Title:       spotifyAlbum.Name,
			AlbumType:   spotifyAlbum.AlbumType,
			Monitor:     artist.Monitor,
			ReleaseDate: spotifyAlbum.ReleaseDateTime(),
		}
		album.Store()
	}

	return nil
}
