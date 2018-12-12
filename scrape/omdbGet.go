package movieapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/calshius/go-movie-api/client"
)

// OMDB details we want
type OMDB struct {
	ImdbID string `json:"imdbID"`
}

// Parse the returned data from open movie DB api into the JSON struct
func omdbDetails(body []byte) (*OMDB, error) {
	var s = new(OMDB)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

// Query the omdb api to get the imdb id for passing on so we can retrieve revenue numbers
func FetchOMDBDetails(title string) *OMDB {

	apiKey := os.Getenv("OMDB_APIKEY")

	if apiKey == "" {
		log.Printf("Error: OMDB_APIKEY has not been set")
	}

	path, err := url.Parse("http://www.omdbapi.com/")
	if err != nil {
		log.Fatal(err)
	}

	query := path.Query()

	query.Add("apikey", apiKey)
	query.Add("t", title)
	path.RawQuery = query.Encode()

	response := movieapi.HTTPGet(path)

	omdb, err := omdbDetails([]byte(response))

	return omdb
}
