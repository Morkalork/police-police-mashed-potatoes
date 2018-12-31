package main

import (
	"fmt"
	"github.com/police-police-mashed-potatoes/business"
	"github.com/police-police-mashed-potatoes/console"
	"github.com/police-police-mashed-potatoes/data"
)

func main() {
	fmt.Println("Contacting API, this could take some time...")
	fmt.Println("")
	eventInformation := data.LoadFromAPI()
	result := business.ProcessData(eventInformation)
	console.OutputResult(result)
}
