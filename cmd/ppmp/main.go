package main

import (
	"fmt"
	"github.com/police-police-mashed-potatoes/loader"
	"github.com/police-police-mashed-potatoes/processing"
)

func output(data processing.DataDescription) {
	fmt.Println("Current situation in sweden:")
	fmt.Println("_____________________")
	fmt.Printf("Number of events: %d\n", data.Count)
	fmt.Printf("Most common crime type: %s\n", data.MostCommonCrimeType)
	fmt.Printf("Worst location: %s\n", data.WorstLocation.Name)
}

func main() {
	fmt.Println("Contacting API, this could take some time...")
	fmt.Println("")
	data := loader.LoadData()
	result := processing.ProcessData(data)
	output(result)
}
