package service

import (
	"SQLConnection/internal/db"
	"SQLConnection/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAlbums(c *gin.Context) {
	db.GetAllAlbumsFromDb(c)
}

func GetAlbumByTitle(c *gin.Context, title string) {
	db.GetAlbumFromDb(c, title)
}

func PostAlbum(c *gin.Context) {
	var input model.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album := model.Album{
		Title:  input.Title,
		Artist: input.Artist,
		Price:  input.Price,
	}

	db.PostAlbumToDb(c, &album)
	return
}
