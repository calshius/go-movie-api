package main

import (
	"strings"
	"github.com/calshius/go-movie-api/scrape"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/movie", func(c *gin.Context) {

		id := c.Query("name")

		id = strings.Replace(id, " ", "+", -1)

		omdbDetails := movieapi.FetchOMDBDetails(id)

		tmdbResult := movieapi.FetchMovieDetails(omdbDetails.ImdbID)

		c.JSON(200, tmdbResult)
	})

	router.Run(":8080")
}
