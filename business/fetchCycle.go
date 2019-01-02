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
	fmt.Println("Thank you for using Police Police Mashed Potatoes!")
	fmt.Println("Initiating scan...")
	fmt.Println("")

	entries := data.LoadFromAPI(
		config.Hours,
		config.Location,
		config.CrimeType,
	)

	console.OutputResult(entries)
}
