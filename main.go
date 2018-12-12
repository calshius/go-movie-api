package main

import (
	"fmt"
	"github.com/calshius/go-movie-api/scrape"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/movie", func(c *gin.Context) {

		id := c.Query("name")

		omdbDetails := movieapi.FetchOMDBDetails(id)

		tmdbResult := movieapi.FetchMovieDetails(omdbDetails.ImdbID)

		fmt.Println(tmdbResult)

		c.JSON(200, tmdbResult)
	})

	router.Run(":8080")
}
