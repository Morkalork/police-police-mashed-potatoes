package business

import (
	"flag"
)

type EventsConfig struct {
	Hours int
	Location string
	CrimeType string
	MaxRecords int
	ListCrimeTypes bool
	ShowHelp bool
}

func GetArguments() EventsConfig {
	hours := flag.Int("hours", 24, "The number of hours to back-check")
	maxRecords := flag.Int("maxrecords", 100000, "The maximum number of records to store")
	location := flag.String("location", "", "The location to monitor, such as Stockholm")
	crimeType := flag.String("crimetype", "", "The crime type to check specifically for")
	listCrimeTypes := flag.Bool("listcrimetypes", false, "Chose whether or not to list all available crime types")
	showHelp := flag.Bool("help", false, "Show the help section")
	flag.Parse()
	return EventsConfig{
		Hours: *hours,
		MaxRecords: *maxRecords,
		Location: *location,
		CrimeType: *crimeType,
		ListCrimeTypes: *listCrimeTypes,
		ShowHelp: *showHelp,
	}
}