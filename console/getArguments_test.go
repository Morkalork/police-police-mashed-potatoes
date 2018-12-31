package console

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type MockUserInput struct{}

func (s *MockUserInput) GetArgument(name string, defaultValue int) UserArgument {
	return UserArgument{
		Name:  name,
		Value: 1,
	}
}

func TestGetArguments_ReturnsArguments(t *testing.T) {
	userInput := &MockUserInput{}
	result := GetArguments(userInput)

	require.True(t, len(result) > 0)
}

func TestGetArguments_ContainsHoursArgument(t *testing.T) {
	result := GetArguments(&MockUserInput{})

	var correctArgument UserArgument
	for _, userArgument := range result {
		if userArgument.Name == "hours" {
			correctArgument = userArgument
		}
	}

	require.Equal(t, 1, correctArgument.Value)
}

func TestGetArguments_ContainsMaxRecordsArgument(t *testing.T) {
	result := GetArguments(&MockUserInput{})

	var correctArgument UserArgument
	for _, userArgument := range result {
		if userArgument.Name == "maxrecords" {
			correctArgument = userArgument
		}
	}

	require.Equal(t, 1, correctArgument.Value)
}
