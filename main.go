package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/movie", func(c *gin.Context) {

		id := c.Query("name")

		path := "http://www.omdbapi.com/?t=" + id + "&apikey=BanMePlz"

		fmt.Println(path)

		resp := fetchMovieDetails(path)

		fmt.Println(resp)

		c.JSON(200, resp)
	})

	router.Run(":8080")
}
