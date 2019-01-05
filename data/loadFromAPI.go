package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	GPS  string `json:"gps"`
}

type Entry struct {
	Id        int    `json:"id"`
	DateTime  string   `json:"datetime"`
	Name      string   `json:"name"`
	Summary   string   `json:"summary"`
	Url       string   `json:"url"`
	CrimeType string   `json:"type"`
	Location  Location `json:"location"`
}

func handleError(err error) {
	log.Fatal("Fatal error: ", err)
}

// loadData will load the loader from the police and return it.
// If anything goes wrong, it will log the error.
func LoadFromAPI(hours int, location string, crimeType string) []Entry {
	var entries []Entry
	hour := 0
	for hour < hours {
		url, urlError := buildUrl(hour, location, crimeType)

		if hour == 0 {
			fmt.Print("Looking at current events (first time can take a bit longer)... ")
		} else {
			plural := ""
			if hour > 1 {
				plural = "s"
			}
			fmt.Printf("Looking at events %d hour%s back... ", hour, plural)
		}

		if urlError != nil {
			handleError(urlError)
			return entries
		}

		resp, httpError := http.Get(url)

		if httpError != nil {
			handleError(httpError)
			return entries
		}


		body, readError := ioutil.ReadAll(resp.Body)

		if readError != nil {
			handleError(readError)
			return entries
		}

		var foundEntries []Entry
		jsonError := json.Unmarshal(body, &foundEntries)
		entries = append(entries, foundEntries...)

		if jsonError != nil {
			handleError(jsonError)
			return entries
		}

		responseBodyError := resp.Body.Close()

		if responseBodyError != nil {
			handleError(responseBodyError)
			return entries
		}

		fmt.Printf("found %d events\n", len(foundEntries))
		hour++
	}

	return entries
}
