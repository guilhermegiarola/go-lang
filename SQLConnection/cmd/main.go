package main

import (
	"SQLConnection/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	//initializes the gin router and uses the GET function to
	//associate the method to the path
	router := gin.Default()

	router.Handle("GET", "/albums", func(c *gin.Context) {
		title := c.Query("title")
		if title != "" {
			controllers.GetAlbumByTitle(c)
		} else {
			controllers.GetAlbums(c)
		}
	})

	//router.GET("/albums", controllers.GetAlbums)

	//uses the POST function also to associate the
	//post method to the path
	router.POST("/albums", controllers.PostAlbum)

	//attaches the router to an http.Server and starts the server
	router.Run("localhost:8080")
}
