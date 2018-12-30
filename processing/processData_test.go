package processing

import (
	"../../police-police-mashed-potatoes/loader"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProcessData_EmptyData_ReturnsEmptySlice(test *testing.T) {
	result := ProcessData([]loader.Entry{})
	require.Empty(test, result)
}

func TestProcessData_WithEntries_CanReturnProperCount(test *testing.T) {
	data := []loader.Entry{
		{Id: 1},
		{Id: 2},
	}

	result := ProcessData(data)
	require.True(test, result.Count == 2)
}

func TestProcessData_WithEntries_CanReturnCorrectCrimeType(test *testing.T) {
	data := []loader.Entry{
		{CrimeType: "Theft"},
		{CrimeType: "Assault"},
		{CrimeType: "Theft"},
	}

	result := ProcessData(data)
	require.Equal(test, result.MostCommonCrimeType, "Theft")
}

func TestProcessData_WithEntries_CanReturnCorrectLocation(test *testing.T) {
	data := []loader.Entry{
		{Location: loader.Location{Name: "Lund"} },
		{Location: loader.Location{Name: "Malmö"}},
		{Location: loader.Location{Name: "Malmö"}},
	}

	result := ProcessData(data)
	require.Equal(test, result.WorstLocation.Name, "Malmö")
}
