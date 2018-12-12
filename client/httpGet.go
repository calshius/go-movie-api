package movieapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// HTTPGet Preset HTTP settings for client gets
func HTTPGet(url *url.URL) []byte {

	httpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

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

	return body
}
