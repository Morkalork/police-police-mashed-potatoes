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
	require.Equal(t, 1, config.Hours)
	require.Equal(t, 2, config.MaxRecords)
	require.Equal(t, "Lund", config.Location)
	require.Equal(t, "Bombhot", config.CrimeType)
}
