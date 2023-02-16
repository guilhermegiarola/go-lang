package db

import (
	"SQLConnection/internal/model"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func OpenConnection() *mongo.Collection {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client.Database("testing").Collection("Albums")
}

func PutAlbumToDb(album *model.Album) {
	var collection = OpenConnection()
	collection.InsertOne(context.TODO(), album)
}

func GetAllAlbumsFromDb(c *gin.Context) []model.Album {
	var collection = OpenConnection()
	var albums []model.Album
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		panic(err)
	}
	cursor.All(c, bson.M{})
	for cursor.TryNext(c) {
		var result model.Album
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		albums = append(albums, result)
	}
	fmt.Println(albums)
	return albums
}
