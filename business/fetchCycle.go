package business

import (
	"fmt"
	"github.com/police-police-mashed-potatoes/console"
	"github.com/police-police-mashed-potatoes/data"
)

type FetchCycle struct {
	hours     int
	location  string
	crimeType string
}

func (f FetchCycle) Start(config EventsConfig) {
	fmt.Println("Launching!")

	entries := data.LoadFromAPI(
		config.Hours,
		config.Location,
		config.CrimeType,
	)

	fmt.Println("Done!")
	fmt.Println("")

	console.OutputResult(entries)
}
