package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	//binds the request body to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	//appends the new album to the slice and adds
	// a 201 status code to the response
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusFound, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	//initializes the gin router and uses the GET function to
	//associate the method to the path
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumByID)

	//uses the POST function also to associate the
	//post method to the path
	router.POST("/albums", postAlbums)

	//attaches the router to an http.Server and starts the server
	router.Run("localhost:8080")
}
