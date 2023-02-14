package controllers

import (
	"SQLConnection/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album

	//binds the request body to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//appends the new album to the slice and adds
	// a 201 status code to the response
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusFound, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
