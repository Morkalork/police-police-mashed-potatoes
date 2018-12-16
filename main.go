package main

import (
	"fmt"
)

func output(data DataDescription) {
	fmt.Println("OUTPUT:")
	fmt.Println("_____________________")
	fmt.Printf("Number of events: %d\n", data.Count)
	fmt.Printf("Most common crime type: %s\n", data.MostCommonCrimeType)
	fmt.Printf("Worst location: %s\n", data.WorstLocation.Name)
}

func main() {
	data := loadData()
	result := processData(data)
	output(result)
}
