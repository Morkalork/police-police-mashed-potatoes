package console

import (
	"fmt"
	"github.com/police-police-mashed-potatoes/data"
	"sort"
	"strings"
)

func OutputResult(entries []data.Entry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].CrimeType < entries[j].CrimeType
	})

	fmt.Println("")
	fmt.Printf("Found the following %d event(s):\n", len(entries))
	currentCrimeType := ""
	for _, entry := range entries {
		if currentCrimeType != entry.CrimeType {
			currentCrimeType = entry.CrimeType
			fmt.Println(strings.ToUpper(currentCrimeType))
		}

		fmt.Printf(
			"   %s: (%s) %s\n",
			entry.DateTime,
			entry.Location.Name,
			entry.Summary,
		)
	}
}
