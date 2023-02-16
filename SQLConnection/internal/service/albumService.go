package service

import (
	"SQLConnection/internal/db"
	"SQLConnection/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllAlbums(c *gin.Context) []model.Album {
	return db.GetAllAlbumsFromDb(c)
}

func PutAlbum(c *gin.Context) {
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

	db.PutAlbumToDb(&album)
	return
}
