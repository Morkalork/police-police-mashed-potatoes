package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessData_EmptyData_ReturnsEmptySlice(test *testing.T) {
	result := processData([]Entry{})
	require.Empty(test, result)
}

func TestProcessData_WithEntries_CanReturnProperCount(test *testing.T) {
	data := []Entry{
		{Id: 1},
		{Id: 2},
	}

	result := processData(data)
	require.True(test, result.Count == 2)
}

func TestProcessData_WithEntries_CanReturnCorrectCrimeType(test *testing.T) {
	data := []Entry{
		{CrimeType: "Theft"},
		{CrimeType: "Assault"},
		{CrimeType: "Theft"},
	}

	result := processData(data)
	require.Equal(test, result.MostCommonCrimeType, "Theft")
}

func TestProcessData_WithEntries_CanReturnCorrectLocation(test *testing.T) {
	data := []Entry{
		{Location: Location{Name: "Lund"} },
		{Location: Location{Name: "Malmö"}},
		{Location: Location{Name: "Malmö"}},
	}

	result := processData(data)
	require.Equal(test, result.WorstLocation.Name, "Malmö")
}
