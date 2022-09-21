package controllers

import (
	"net/http"

	"gofy/responses"
	spotifyInfra "gofy/spotify"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify/v2"
)

type SearchController struct{}

func (h SearchController) SearchArtist(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	client, ctx := spotifyInfra.GetClient()
	result, err := client.Search(ctx, name, spotify.SearchTypeArtist)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	artists := make([]responses.SimpleArtist, len(result.Artists.Artists))

	for i, artist := range result.Artists.Artists {
		thumbnails := make([]responses.Thumbnail, len(artist.Images))

		for j, image := range artist.Images {
			thumbnails[j] = responses.Thumbnail{
				Url:    image.URL,
				Height: image.Height,
				Width:  image.Width,
			}
		}

		artists[i] = responses.SimpleArtist{
			Id:         string(artist.ID),
			Name:       artist.Name,
			Thumbnails: thumbnails,
		}
	}

	c.JSON(http.StatusOK, artists)
}

func (h SearchController) SearchAlbum(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	client, ctx := spotifyInfra.GetClient()
	result, err := client.Search(ctx, title, spotify.SearchTypeAlbum)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	albums := make([]responses.SimpleAlbum, len(result.Albums.Albums))

	for i, album := range result.Albums.Albums {
		thumbnails := make([]responses.Thumbnail, len(album.Images))

		for j, image := range album.Images {
			thumbnails[j] = responses.Thumbnail{
				Url:    image.URL,
				Height: image.Height,
				Width:  image.Width,
			}
		}

		albums[i] = responses.SimpleAlbum{
			Id:         string(album.ID),
			Title:      album.Name,
			Thumbnails: thumbnails,
		}
	}

	c.JSON(http.StatusOK, albums)
}

func (h SearchController) SearchTrack(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	client, ctx := spotifyInfra.GetClient()
	result, err := client.Search(ctx, title, spotify.SearchTypeTrack)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	tracks := make([]responses.SimpleTrack, len(result.Tracks.Tracks))

	for i, track := range result.Tracks.Tracks {
		thumbnails := make([]responses.Thumbnail, len(track.Album.Images))

		for j, image := range track.Album.Images {
			thumbnails[j] = responses.Thumbnail{
				Url:    image.URL,
				Height: image.Height,
				Width:  image.Width,
			}
		}

		tracks[i] = responses.SimpleTrack{
			Id:         string(track.ID),
			Title:      track.Name,
			Thumbnails: thumbnails,
		}
	}

	c.JSON(http.StatusOK, tracks)
}

func (h SearchController) SearchPlaylist(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	client, ctx := spotifyInfra.GetClient()
	result, err := client.Search(ctx, title, spotify.SearchTypePlaylist)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		c.Abort()
		return
	}

	playlists := make([]responses.SimplePlaylist, len(result.Playlists.Playlists))

	for i, playlist := range result.Playlists.Playlists {
		thumbnails := make([]responses.Thumbnail, len(playlist.Images))

		for j, image := range playlist.Images {
			thumbnails[j] = responses.Thumbnail{
				Url:    image.URL,
				Height: image.Height,
				Width:  image.Width,
			}
		}

		playlists[i] = responses.SimplePlaylist{
			Id:         string(playlist.ID),
			Title:      playlist.Name,
			Thumbnails: thumbnails,
		}
	}

	c.JSON(http.StatusOK, playlists)
}
