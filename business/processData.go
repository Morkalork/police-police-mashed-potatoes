package business

import "github.com/police-police-mashed-potatoes/data"

type DataDescription struct {
	Count               int
	WorstLocation       data.Location
	MostCommonCrimeType string
}

func processCrimeType(entries []data.Entry) string {
	var crimeTypes = make(map[string]int)

	for _, entry := range entries {
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

func processLocation(entries []data.Entry) data.Location {
	var locations = make(map[data.Location]int)

	for _, entry := range entries {
		if _, ok := locations[entry.Location]; !ok {
			locations[entry.Location] = 0
		}

		locations[entry.Location] += 1
	}

	mostCommonLocation := data.Location{}
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
func ProcessData(entries []data.Entry) DataDescription {
	dataDescription := DataDescription{}
	dataDescription.Count = len(entries)
	dataDescription.MostCommonCrimeType = processCrimeType(entries)
	dataDescription.WorstLocation = processLocation(entries)

	return dataDescription
}
