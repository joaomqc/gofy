package server

import (
	"gofy/config"
	"gofy/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	c := config.GetConfig()
	cacheDir := c.GetString("cache.directory")

	router := gin.Default()
	router.Use(cors.Default())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	searchGroup := router.Group("search")
	{
		search := new(controllers.SearchController)
		searchGroup.GET("artist/:name", search.SearchArtist)
		searchGroup.GET("album/:title", search.SearchAlbum)
		searchGroup.GET("track/:title", search.SearchTrack)
		searchGroup.GET("playlist/:title", search.SearchPlaylist)
	}

	artistGroup := router.Group("artist")
	{
		artist := new(controllers.ArtistController)
		artistGroup.POST("/", artist.AddArtist)
		artistGroup.GET("/", artist.GetArtists)
	}

	router.Static("/image", cacheDir)
	return router
}
