package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/url"
)

func main() {
	router := gin.Default()

	router.POST("/movie", func(c *gin.Context) {

		id := c.Query("name")

		path, err := url.Parse("http://www.omdbapi.com/")
		
		if err != nil {
			log.Fatal(err)
		}

		query := path.Query()

		query.Add("apikey", "BanMePlz")
		query.Add("t", id)
		path.RawQuery = query.Encode()

		// ?t=" + id + "&apikey=BanMePlz"

		fmt.Println(path)

		resp := fetchMovieDetails(path)

		fmt.Println(resp)

		c.JSON(200, resp)
	})

	router.Run(":8080")
}
