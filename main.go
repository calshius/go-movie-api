package main

import (
	"fmt"
	"strings"
	"github.com/calshius/go-movie-api/scrape"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/movie", func(c *gin.Context) {

		id := c.Query("name")

		id = strings.Replace(id, " ", "+", -1)

		omdbDetails := movieapi.FetchOMDBDetails(id)

		tmdbResult := movieapi.FetchMovieDetails(omdbDetails.ImdbID)

		fmt.Println(tmdbResult)

		c.JSON(200, tmdbResult)
	})

	router.Run(":8080")
}
