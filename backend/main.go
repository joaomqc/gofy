package main

import (
	"flag"
	"fmt"
	"os"

	"gofy/config"
	"gofy/db"
	"gofy/mq"
	"gofy/server"
	"gofy/spotify"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	spotify.Init()
	mq.Init()
	server.Init()
}

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/zmb3/spotify/v2"
// )

// func main() {
// 	r := gin.Default()

// 	r.GET("/artist/:id", getArtist)
// 	r.GET("/album/:id", getAlbum)
// 	r.GET("/playlist/:id", getPlaylist)

// 	r.GET("/search/artist/:name", searchArtist)
// 	r.GET("/search/album/:title", searchAlbum)
// 	r.GET("/search/playlist/:title", searchPlaylist)

// 	r.Run()
// }

// func getArtist(c *gin.Context) {
// 	id := (spotify.ID)(c.Param("id"))
// 	client, ctx := GetClient()
// 	result, err := client.GetArtist(ctx, id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// func getAlbum(c *gin.Context) {
// 	id := (spotify.ID)(c.Param("id"))
// 	client, ctx := GetClient()
// 	result, err := client.GetAlbum(ctx, id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// func getPlaylist(c *gin.Context) {
// 	id := (spotify.ID)(c.Param("id"))
// 	client, ctx := GetClient()
// 	result, err := client.GetPlaylist(ctx, id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, result)
// }

// func searchArtist(c *gin.Context) {
// 	name := c.Param("name")
// 	client, ctx := GetClient()
// 	result, err := client.Search(ctx, name, spotify.SearchTypeArtist)

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, result.Artists)
// }

// func searchAlbum(c *gin.Context) {
// 	name := c.Param("title")
// 	client, ctx := GetClient()
// 	result, err := client.Search(ctx, name, spotify.SearchTypeAlbum)

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, result.Artists)
// }

// func searchPlaylist(c *gin.Context) {
// 	name := c.Param("title")
// 	client, ctx := GetClient()
// 	result, err := client.Search(ctx, name, spotify.SearchTypePlaylist)

// 	if err != nil {
// 		panic(err)
// 	}

// 	c.JSON(http.StatusOK, result.Artists)
// }
