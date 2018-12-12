package movieapi

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/url"
	"os"

	"github.com/calshius/go-movie-api/client"
	"github.com/negah/percent"
)

// Movie details we want
type Movie struct {
	Name        string `json:"original_title"`
	Budget      int    `json:"budget"`
	Revenue     int    `json:"revenue"`
	ReleaseDate string `json:"release_date"`
}

// MovieWithPercentages will store a complete returned object
type MovieWithPercentages Movie

// MarshalJSON to add fields dynamically to the final struct
func (m Movie) MarshalJSON() ([]byte, error) {
	// Determine movie success
	if m.Revenue > m.Budget {
		// If the percenatge is greater than 100 calculate profits
		p := percent.PercentOf(m.Revenue, m.Budget)
		p = math.Floor(p)
		e := m.Revenue - m.Budget
		return json.Marshal(struct {
			MovieWithPercentages
			Percentage float64 
			Profit     int
		}{
			MovieWithPercentages: MovieWithPercentages(m),
			Percentage:           p,
			Profit:               e,
		})
	} else {
		// If the percenatge is less than 100 calculate loses
		p := percent.PercentOf(m.Budget, m.Revenue)
		p = math.Floor(p)
		l := m.Budget - m.Revenue
		return json.Marshal(struct {
			MovieWithPercentages
			Percentage float64
			Loses      int
		}{
			MovieWithPercentages: MovieWithPercentages(m),
			Percentage:           p,
			Loses:                l,
		})
	}

}

// Parse the returned data from The movie DB api into the JSON struct
func movieDetails(body []byte) (*Movie, error) {
	var s = new(Movie)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// Query the The moveie DB API with the imdb received from the OMDB function
func FetchMovieDetails(imdbID string) *Movie {

	apiKey := os.Getenv("TMDB_APIKEY")

	if apiKey == "" {
		log.Printf("Error: TMDB_APIKEY has not been set")
	}

	path, err := url.Parse("https://api.themoviedb.org/3/movie/" + imdbID)
	if err != nil {
		log.Fatal(err)
	}

	query := path.Query()

	query.Add("api_key", apiKey)
	query.Add("language", "en-US")
	path.RawQuery = query.Encode()

	response := movieapi.HTTPGet(path)

	movie, err := movieDetails([]byte(response))

	return movie
}
