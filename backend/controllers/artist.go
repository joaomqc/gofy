package controllers

import (
	"net/http"

	"gofy/forms"
	"gofy/handlers"

	"github.com/gin-gonic/gin"
)

type ArtistController struct{}

var artistHandler = new(handlers.ArtistHandler)

func (a ArtistController) AddArtist(c *gin.Context) {
	var form forms.AddArtist
	if err := c.ShouldBindJSON(&form); err != nil {
		panic(err)
	}

	if form.SpotifyId == "" {
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	artist, err := artistHandler.Add(form)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, artist)
}

func (a ArtistController) GetArtists(c *gin.Context) {
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	host := scheme + "://" + c.Request.Host
	artists, err := artistHandler.GetAll(host)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, artists)
}
