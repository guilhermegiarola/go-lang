package main

import (
	"SQLConnection/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	//initializes the gin router and uses the GET function to
	//associate the method to the path
	router := gin.Default()

	router.GET("/albums", controllers.GetAlbums)

	// router.GET("/albums/:id", controllers.GetAlbumByID)

	//uses the PUT function also to associate the
	//put method to the path
	router.PUT("/albums", controllers.PutAlbum)

	//attaches the router to an http.Server and starts the server
	router.Run("localhost:8080")
}
