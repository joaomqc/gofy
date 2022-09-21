package handlers

import (
	"fmt"
	"gofy/config"
	"gofy/models"
	spotifyInfra "gofy/spotify"
	"os/exec"

	"github.com/zmb3/spotify/v2"
)

type TrackHandler struct{}

func (h TrackHandler) LoadFromAlbum(album models.Album) error {
	client, ctx := spotifyInfra.GetClient()
	tracksPage, err := client.GetAlbumTracks(ctx, album.SpotifyID, spotify.Market("PT"))

	if err != nil {
		return err
	}

	spotifyTracks := tracksPage.Tracks

	for {
		if tracksPage.Next == "" {
			break
		}

		err = client.NextPage(ctx, tracksPage)
		if err != nil {
			return err
		}
		spotifyTracks = append(spotifyTracks, tracksPage.Tracks...)
	}

	for _, spotifyTrack := range spotifyTracks {
		track := models.Track{
			SpotifyID: spotifyTrack.ID,
			AlbumID:   album.ID,
			Title:     spotifyTrack.Name,
			Monitor:   album.Monitor,
		}
		track.Store()
	}

	return nil
}

func (h TrackHandler) DownloadTrack(track models.Track) error {
	c := config.GetConfig()
	musicDirectory := c.GetString("download.directory")

	app := "spotdl"
	outDirOption := "--output-directory" + musicDirectory
	url := "https://open.spotify.com/track/" + track.SpotifyID.String()

	cmd := exec.Command(app, outDirOption, url)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
