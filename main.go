package main

//At the command line, use go get to add the github.com/gin-gonic/gin module as a dependency for your module. 
//Use a dot argument to mean “get dependencies for code in the current directory.”
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

// gin.Context contains requests details, validates and serializes JSON and more. 
func getAlbums (c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums) // serializes the struct into JSON and add it into the response. 
}

func postAlbums (c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect album format/syntax."})
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById (c *gin.Context){
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default() //Initialize a Gin router using Default.
    router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("albums", postAlbums)
    router.Run("localhost:8080") //Use the Run function to attach the router to an http.Server and start the server.
}
