package spotify

import (
	"context"
	"gofy/config"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

var client *spotify.Client
var ctx context.Context

func Init() {
	ctx = context.Background()
	c := config.GetConfig()
	config := &clientcredentials.Config{
		ClientID:     c.GetString("spotify.id"),
		ClientSecret: c.GetString("spotify.secret"),
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(ctx)
	if err != nil {
		panic(err)
	}
	httpClient := spotifyauth.New().Client(ctx, token)
	client = spotify.New(httpClient)
}

func GetClient() (*spotify.Client, context.Context) {
	return client, ctx
}
