package business

import (
	"github.com/police-police-mashed-potatoes/data"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessData_EmptyData_ReturnsEmptySlice(test *testing.T) {
	result := ProcessData([]data.Entry{})
	require.Empty(test, result)
}

func TestProcessData_WithEntries_CanReturnProperCount(test *testing.T) {
	data := []data.Entry{
		{Id: 1},
		{Id: 2},
	}

	result := ProcessData(data)
	require.True(test, result.Count == 2)
}

func TestProcessData_WithEntries_CanReturnCorrectCrimeType(test *testing.T) {
	data := []data.Entry{
		{CrimeType: "Theft"},
		{CrimeType: "Assault"},
		{CrimeType: "Theft"},
	}

	result := ProcessData(data)
	require.Equal(test, result.MostCommonCrimeType, "Theft")
}

func TestProcessData_WithEntries_CanReturnCorrectLocation(test *testing.T) {
	data := []data.Entry{
		{Location: data.Location{Name: "Lund"} },
		{Location: data.Location{Name: "Malmö"}},
		{Location: data.Location{Name: "Malmö"}},
	}

	result := ProcessData(data)
	require.Equal(test, result.WorstLocation.Name, "Malmö")
}
