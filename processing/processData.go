package processing

import (
	"github.com/police-police-mashed-potatoes/loader"
)

type DataDescription struct {
	Count               int
	WorstLocation       loader.Location
	MostCommonCrimeType string
}

func processCrimeType(data []loader.Entry) string {
	var crimeTypes = make(map[string]int)

	for _, entry := range data {
		if _, ok := crimeTypes[entry.CrimeType]; !ok {
			crimeTypes[entry.CrimeType] = 0
		}
		crimeTypes[entry.CrimeType] += 1
	}

	mostCommonCrimeTime := ""
	highestCount := 0
	for key, val := range crimeTypes {
		if val > highestCount {
			highestCount = val
			mostCommonCrimeTime = key
		}
	}

	return mostCommonCrimeTime
}

func processLocation(data []loader.Entry) loader.Location {
	var locations = make(map[loader.Location]int)

	for _, entry := range data {
		if _, ok := locations[entry.Location]; !ok {
			locations[entry.Location] = 0
		}

		locations[entry.Location] += 1
	}

	mostCommonLocation := loader.Location{}
	highestCount := 0
	for key, val := range locations {
		if val > highestCount {
			highestCount = val
			mostCommonLocation = key
		}
	}

	return mostCommonLocation
}

// processData will create a struct with information describing the loader
func ProcessData(data []loader.Entry) DataDescription {
	dataDescription := DataDescription{}
	dataDescription.Count = len(data)
	dataDescription.MostCommonCrimeType = processCrimeType(data)
	dataDescription.WorstLocation = processLocation(data)

	return dataDescription
}
