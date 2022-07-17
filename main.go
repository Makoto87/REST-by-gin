package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
// gin.Contextはリクエストを検証し、JSONにシリアライズしてくれる
func getAlbums(c *gin.Context) {
	// IndentedJSON: 構造体をJSONにシリアライズし、レスポンスとして返す
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	// router の初期化
	router := gin.Default()
	// GETメソッドと、/albums パスを関連づける
	router.GET("/albums", getAlbums)

	// routerとhttpサーバーを紐付け、サーバーを立ち上げる
	router.Run("localhost:8080")
}
