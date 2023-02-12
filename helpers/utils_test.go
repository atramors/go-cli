package helpers

import (
	"testing"
)

type TestCase struct {
	Name     string
	Input    string
	Expected bool
}

func TestArgIsNum(t *testing.T) {
	testMapper := []TestCase{
		{
			Name:     "Positive with only digits",
			Input:    "777",
			Expected: true,
		},
		{
			Name:     "Negative with weird symbol",
			Input:    "৩",
			Expected: false,
		},
		{
			Name:     "Negative with only letters",
			Input:    "SomeString",
			Expected: false,
		},
		{
			Name:     "Negative with letters and digits",
			Input:    "Some777String",
			Expected: false,
		},
	}
	for _, test := range testMapper {
		t.Run(test.Name, func(t *testing.T) {
			result := ArgIsNum(test.Input)
			if result != test.Expected {
				t.Errorf("%t != %t with input: %s", result, test.Expected, test.Input)
			}
		})
	}
}
