package loader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
	GPS string `json:"gps"`
}

type Entry struct {
	Id int64 `json:"id"`
	DateTime string `json:"datetime"`
	Name string `json:"name"`
	Summary string `json:"summary"`
	Url string `json:"url"`
	CrimeType string `json:"type"`
	Location Location `json:"location"`
}

func handleError(err error) {
	log.Fatal("Fatal error: ", err)
}

// loadData will load the loader from the police and return it.
// If anything goes wrong, it will log the error.
func LoadData() []Entry {
	var entries []Entry
	url := "https://polisen.se/api/events"

	resp, err := http.Get(url)

	if err != nil {
		handleError(err)
		return entries
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		handleError(err)
		return entries
	}

	err = json.Unmarshal(body, &entries)

	if err != nil {
		handleError(err)
		return entries
	}

	return entries
}