package main

import (
	"fmt"
	"github.com/police-police-mashed-potatoes/business"
	"github.com/police-police-mashed-potatoes/data"
)

func showHelp() {
	fmt.Println("HELP - Police Police Mashed Potatoes")
	fmt.Println("")
	fmt.Println("This app is used to generate a file of statistics based on the Swedetish police events API.")
	fmt.Println("You can look back up to roughly 48 hours (possibly more if the police feels like it). Try more if you want to live dangerously.")
	fmt.Println("It will generate a file called stats.txt with the combined information.")
	fmt.Println("")
	fmt.Println("The following flags are available:")
	fmt.Println("-hours=X          Set how many hours back to check, defaults to 24")
	fmt.Println("-maxrecords=X     How many records to save, will stop processing if this is reached. Defaults to 100.000")
	fmt.Println("-location=X       Set which location to look in, such as Stockholm. Leave empty for all")
	fmt.Println("-crimetype=X      Which crime type to look for. Leave empty for all")
	fmt.Println("-listcrimetypes   Show all available crime types")
	fmt.Println("-help             Show this help section")
}

func showCrimeTypes() {
	for _, crimeType := range data.CrimeTypes {
		fmt.Printf("%s\n", crimeType)
	}
}

func main() {
	config := business.GetArguments()
	if config.ShowHelp {
		showHelp()
	} else if config.ListCrimeTypes {
		showCrimeTypes()
	} else {
		cycle := business.FetchCycle{}
		cycle.Start(config)
	}
}
