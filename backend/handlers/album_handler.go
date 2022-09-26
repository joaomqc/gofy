package handlers

import (
	"fmt"
	"gofy/config"
	"gofy/models"
	spotifyInfra "gofy/spotify"
	"os"

	"github.com/zmb3/spotify/v2"
)

type AlbumHandler struct{}

var albumTypes = []spotify.AlbumType{
	spotify.AlbumTypeAlbum,
	spotify.AlbumTypeSingle,
}

func (h AlbumHandler) LoadFromArtist(artist models.Artist) error {
	c := config.GetConfig()
	cacheDir := c.GetString("cache.directory")

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

		albumCacheDir := fmt.Sprintf("%s/albums/%d", cacheDir, album.ID)
		os.MkdirAll(albumCacheDir, os.ModePerm)
		for _, image := range spotifyAlbum.Images {
			filename := fmt.Sprintf("%s/%d-%d.jpeg", albumCacheDir, image.Width, image.Height)
			f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				return err
			}
			defer f.Close()
			image.Download(f)
		}
	}

	return nil
}
