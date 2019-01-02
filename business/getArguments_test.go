package business

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func resetter () {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
}

func TestGetArguments(t *testing.T) {
	resetter()
	os.Args = []string{"path", "-hours=1", "-maxrecords=2", "--location=Lund", "--crimetype=Bombhot"}
	config := GetArguments()
	require.Equal(t, 1, config.hours)
	require.Equal(t, 2, config.maxRecords)
	require.Equal(t, "Lund", config.location)
	require.Equal(t, "Bombhot", config.crimeType)
}
