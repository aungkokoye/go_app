package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aungkokoye/go_app/cmd"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {

	fmt.Println("starting go app!")

	// router := gin.Default()
	// router.GET("/albums", getAlbums)

	// router.Run("localhost:8080")

	// fmt.Println("Running go_api_gin!")

	err := cmd.Execute()
	if err != nil {
		log.Println(err)

		os.Exit(1)
	}

	os.Exit(0)

}

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }
