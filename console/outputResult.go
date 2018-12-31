package console

import (
	"fmt"
	"github.com/police-police-mashed-potatoes/business"
)

func OutputResult(data business.DataDescription) {
	fmt.Println("Current situation in sweden:")
	fmt.Println("_____________________")
	fmt.Printf("Number of events: %d\n", data.Count)
	fmt.Printf("Most common crime type: %s\n", data.MostCommonCrimeType)
	fmt.Printf("Worst location: %s\n", data.WorstLocation.Name)
}
