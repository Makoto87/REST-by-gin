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

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	// リクエストボディを newAlbum に紐づける
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	// id パラメータを取得
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	// IDとマッチするalbumを探し、あったらJSONとして返す
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	// IDとマッチするalbumがなかったら404エラーを返す
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// router の初期化
	router := gin.Default()

	// ginはHTTPメソッドとパスをハンドラーと紐づけることが可能
	// GETメソッドと /albums パス、getAlbums 関数を関連づける
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	// POSTメソッドと /albums パス、postAlbums 関数を関連づける
	router.POST("/albums", postAlbums)

	// routerとhttpサーバーを紐付け、サーバーを立ち上げる
	router.Run("localhost:8080")
}
