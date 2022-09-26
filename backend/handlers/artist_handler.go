package handlers

import (
	"fmt"
	"gofy/config"
	"gofy/db"
	"gofy/forms"
	"gofy/models"
	"gofy/responses"
	spotifyInfra "gofy/spotify"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type ArtistHandler struct{}

var albumHandler = new(AlbumHandler)

func (h ArtistHandler) Add(form forms.AddArtist) (*models.Artist, error) {
	c := config.GetConfig()
	cacheDir := c.GetString("cache.directory")

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

	artistCacheDir := fmt.Sprintf("%s/artists/%d", cacheDir, artist.ID)
	os.MkdirAll(artistCacheDir, os.ModePerm)
	for _, image := range result.Images {
		filePath := fmt.Sprintf("%s/%d-%d.jpeg", artistCacheDir, image.Width, image.Height)
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		image.Download(f)
	}

	return artist, nil
}

func (h ArtistHandler) GetAll(host string) ([]responses.SimpleArtist, error) {
	c := config.GetConfig()
	cacheDir := c.GetString("cache.directory")

	sqlStatement := `
		SELECT id, name, date_added
		FROM artist
	`

	rows, err := db.GetDB().Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	artists := []responses.SimpleArtist{}

	for rows.Next() {
		artist := models.Artist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.DateAdded)
		if err != nil {
			return nil, err
		}
		artistCacheDir := fmt.Sprintf("%s/artists/%d", cacheDir, artist.ID)
		files, err := ioutil.ReadDir(artistCacheDir)
		if err != nil {
			return nil, err
		}
		thumbnails := make([]responses.Thumbnail, len(files))

		for i, file := range files {
			splitByExt := strings.Split(file.Name(), ".")
			splitName := strings.Split(splitByExt[0], "-")
			width, err := strconv.Atoi(splitName[0])
			if err != nil {
				return nil, err
			}
			height, err := strconv.Atoi(splitName[1])
			if err != nil {
				return nil, err
			}

			thumbnails[i] = responses.Thumbnail{
				Url:    fmt.Sprintf("%s/image/artists/%d/%s", host, artist.ID, file.Name()),
				Height: height,
				Width:  width,
			}
		}

		artists = append(artists, responses.SimpleArtist{
			Id:         fmt.Sprint(artist.ID),
			Name:       artist.Name,
			Thumbnails: thumbnails,
		})
	}

	return artists, nil
}
