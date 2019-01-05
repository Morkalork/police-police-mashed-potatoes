package business

import (
	"github.com/police-police-mashed-potatoes/data"
	"strings"
)

func inArray(arr []int, val int) bool {
	for _,arrVal := range arr {
		if arrVal == val {
			return true
		}
	}

	return false
}

func SanitizeEntries(entries []data.Entry) []data.Entry {
	var result []data.Entry
	var ids []int
	var newLineReplacer = strings.NewReplacer("\n", "")
	for _,entry := range entries {
		if strings.Contains(entry.DateTime, "\n") {
			entry.DateTime = newLineReplacer.Replace(entry.DateTime)
		}

		if !inArray(ids, entry.Id) {
			ids = append(ids, entry.Id)
			result = append(result, entry)
		}
	}

	return result
}
