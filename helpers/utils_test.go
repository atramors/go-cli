package helpers

import (
	"testing"
)

type testCase struct {
	name     string
	input    string
	expected bool
}

func TestArgIsNum(t *testing.T) {
	testMapper := []testCase{
		{
			name:     "Positive with only digits",
			input:    "777",
			expected: true,
		},
		{
			name:     "Negative with weird symbol",
			input:    "৩",
			expected: false,
		},
		{
			name:     "Negative with only letters",
			input:    "SomeString",
			expected: false,
		},
		{
			name:     "Negative with letters and digits",
			input:    "Some777String",
			expected: false,
		},
	}
	for _, test := range testMapper {
		t.Run(test.name, func(t *testing.T) {
			result := ArgIsNum(test.input)
			if result != test.expected {
				t.Errorf("%t != %t with input: %s", result, test.expected, test.input)
			}
		})
	}
}
