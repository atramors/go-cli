package helpers

import (
	"os"
	"testing"
)

type TestCase struct {
	Name         string
	Input        string
	InputSlc     []string
	ExpectedBool bool
	ExpectedStr  string
	ExpectedSlc  []string
}

func TestArgIsNum(t *testing.T) {
	testMapper := []TestCase{
		{
			Name:         "Positive with only digits",
			Input:        "777",
			ExpectedBool: true,
		},
		{
			Name:         "Negative with weird symbol",
			Input:        "৩",
			ExpectedBool: false,
		},
		{
			Name:         "Negative with only letters",
			Input:        "SomeString",
			ExpectedBool: false,
		},
		{
			Name:         "Negative with letters and digits",
			Input:        "Some777String",
			ExpectedBool: false,
		},
	}
	for _, test := range testMapper {
		t.Run(test.Name, func(t *testing.T) {
			result := ArgIsNum(test.Input)
			if result != test.ExpectedBool {
				t.Errorf("'%t' != '%t' with input: '%s'", result, test.ExpectedBool, test.Input)
			}
		})
	}
}

func TestArgProcOne(t *testing.T) {
	testMapper := []TestCase{
		{
			Name:        "Positive with letters",
			Input:       "WonderCity",
			ExpectedStr: "WonderCity",
		},
		{
			Name:        "Positive with empty string",
			Input:       "",
			ExpectedStr: "",
		},
		{
			Name:        "Positive with 0",
			Input:       "0",
			ExpectedStr: "",
		},
	}
	for _, test := range testMapper {
		t.Run(test.Name, func(t *testing.T) {
			os.Args = []string{"cmd", test.Input}
			result := ArgProcessing()
			if result != test.ExpectedStr {
				t.Errorf("'%s' != '%s' with input: '%s'", result, test.ExpectedStr, test.Input)
			}
		})
	}
}

func TestArgProcWithTooManyArgs(t *testing.T) {
	test := TestCase{
		Name:     "Test with more than 1 arg provided",
		InputSlc: []string{"cmd", "WonderCity", "AdditionalArgument"},
	}
	os.Args = test.InputSlc
	result := ArgProcessing()

	if result != test.ExpectedStr {
		t.Errorf("%s failed, expected: '', got:  '%s'", test.Name, result)
	}
}

func TestArgProcWithTooFewArgs(t *testing.T) {
	test := TestCase{
		Name:     "with less than 1 arg povided",
		InputSlc: []string{"cmd"},
	}
	os.Args = test.InputSlc
	result := ArgProcessing()

	if result != test.ExpectedStr {
		t.Errorf("%s failed, expected no args, but got: '%s'", test.Name, result)
	}
}
