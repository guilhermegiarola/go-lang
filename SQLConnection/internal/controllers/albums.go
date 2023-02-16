package controllers

import (
	"SQLConnection/internal/service"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	service.GetAllAlbums(c)
}

func PutAlbum(c *gin.Context) {
	service.PutAlbum(c)
}
