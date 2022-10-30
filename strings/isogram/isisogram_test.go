package isogram

import (
	"testing"
)

var testCases = []struct {
	name     string
	input    string
	expected bool
}{
	{
		"Alphanumeric string 1",
		"copy1",
		false,
	},
	{
		"Alphanumeric string 2",
		"copy1sentence",
		false,
	},
	{
		"Alphanumeric string 3",
		"copy1 sentence with space",
		false,
	},
	{
		"Alphabetic string 1",
		"allowance",
		false,
	},
	{
		"Alphabetic string 2",
		"My Doodle",
		false,
	},
	{
		"Alphabetic string with symbol",
		"Isogram!",
		false,
	},
	{
		"Isogram string 1",
		"Uncopyrightable",
		true,
	},
}

func TestIsIsogram(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			actual := IsIsogram(test.input)

			if test.expected != actual {
				t.Errorf("Expected to be %v for string %s but was %v.", test.expected, test.input, actual)
			}
		})
	}
}
