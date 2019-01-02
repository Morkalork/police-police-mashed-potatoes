package data

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestBuildUrl_WithZeroHours_ReturnsNow(test *testing.T) {
	hours := 0
	location := ""
	crimeType := ""

	resultUrl,_ := buildUrl(hours, location, crimeType)

	now := time.Now()
	nowStr := fmt.Sprintf("%d-%d-%d%%20%d", now.Year(), now.Month(), now.Day(), now.Hour())
	require.True(test, strings.Contains(resultUrl, nowStr))
}
func TestBuildUrl_With25Hours_ReturnsYesterday(test *testing.T) {
	hours := 25
	location := ""
	crimeType := ""

	resultUrl,_ := buildUrl(hours, location, crimeType)

	now := time.Now()
	then := now.AddDate(0, 0, -1)
	yesterdayStr := fmt.Sprintf("%d-%d-%d", then.Year(), then.Month(), then.Day())
	require.True(test, strings.Contains(resultUrl, yesterdayStr))
}

func TestBuildUrl_WithLocation_AddsLocation(test *testing.T) {
	hours := 0
	location := "Sörböle"
	crimeType := ""

	resultUrl,_ := buildUrl(hours, location, crimeType)

	locationStr := fmt.Sprintf("locationname=%s", url.QueryEscape(location))
	require.True(test, strings.Contains(resultUrl, locationStr))
}

func TestBuildUrl_WithValidCrimeType_AddsCrimeType(test *testing.T) {
	hours := 0
	location := ""
	crimeType := "Bombhot"

	resultUrl,_ := buildUrl(hours, location, crimeType)

	crimeTypeStr := fmt.Sprintf("type=%s", url.QueryEscape(crimeType))
	require.True(test, strings.Contains(resultUrl, crimeTypeStr))
}

func TestBuildUrl_WithInvalidCrimeType_ReturnsError(test *testing.T) {
	hours := 0
	location := ""
	crimeType := "Fiffle and baug"

	_,error := buildUrl(hours, location, crimeType)

	require.True(test, error != nil)
}
