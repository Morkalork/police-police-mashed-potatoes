package main

import (
	"testing"
)

func TestProcessData(test *testing.T) {
	test.Run("Can handle empty data", func(t *testing.T) {
		result := processData([]Entry{})
		isEmpty := result == DataDescription{}

		if (!isEmpty) {
			t.Error("Struct was not empty")
		}
	})

	test.Run("Can return a proper count", func(t *testing.T) {
		data := []Entry{
			{Id: 1},
			{Id: 2},
		}

		result := processData(data)
		if result.Count != 2 {
			t.Errorf("Struct has the wrong Count, expected 2 but got %d", result.Count)
		}
	})

	test.Run("Can return the correct crime type", func(t *testing.T) {
		data := []Entry{
			{CrimeType: "Theft"},
			{CrimeType: "Assault"},
			{CrimeType: "Theft"},
		}

		result := processData(data)
		if result.MostCommonCrimeType != "Theft" {
			t.Errorf("Struct has the wrong crime type, expected 'Theft' but got '%s'", result.MostCommonCrimeType)
		}
	})

	test.Run("Can return the correct location", func(t *testing.T) {
		data := []Entry{
			{Location: Location{Name: "Lund"} },
			{Location: Location{Name: "Malmö"}},
			{Location: Location{Name: "Malmö"}},
		}

		result := processData(data)
		if result.WorstLocation.Name != "Malmö" {
			t.Errorf("Struct has the wrong location, expected Malmö but got '%s'", result.WorstLocation.Name)
		}
	})
}
