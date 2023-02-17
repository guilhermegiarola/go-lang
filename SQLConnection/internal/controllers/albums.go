package controllers

import (
	"SQLConnection/internal/service"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	service.GetAllAlbums(c)
}

func GetAlbumByTitle(c *gin.Context) {
	title := c.Query("title")
	service.GetAlbumByTitle(c, title)
}

func PostAlbum(c *gin.Context) {
	service.PostAlbum(c)
}
