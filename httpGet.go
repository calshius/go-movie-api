package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Movie details we want
type Movie struct {
	Name  string `json:"Title"`
	Score string `json:"Metascore"`
}

func movieDetails(body []byte) (*Movie, error) {
	var s = new(Movie)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func fetchMovieDetails(url *url.URL) *Movie {
	httpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	fmt.Println(url)

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	movie, err := movieDetails([]byte(body))

	fmt.Println(movie)

	return movie
}
