package business

import (
	"github.com/police-police-mashed-potatoes/data"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestSanitizeEntries_WithDateTimeNewLines_RemovesNewLines(t *testing.T) {
	entries := []data.Entry{
		{Id: 1, DateTime: "2018-01-01\n"},
	}

	result := SanitizeEntries(entries)
	require.True(t, !strings.Contains(result[0].DateTime, "\n"))
}

func TestSanitizeEntries_WithDuplicateIds_RemovesDuplicates(t *testing.T) {
	entries := []data.Entry{
		{Id: 1},
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 3},
		{Id: 4},
	}

	result := SanitizeEntries(entries)
	require.Len(t, result, 4)
}
