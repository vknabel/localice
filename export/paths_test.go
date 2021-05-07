package export

import (
	"strconv"
	"testing"
)

func TestExpandedPathLocale(t *testing.T) {
	type testCase struct {
		name     string
		path     string
		expected string
	}
	testCases := []testCase{
		{"de", "keep/all", "keep/all"},
		{"en-GB", "keep/all", "keep/all"},
		{"en-US", "keep/all", "keep/all"},

		{"de", "${locale}", "de"},
		{"en-GB", "${locale}", "en-GB"},
		{"en-US", "${locale}", "en-US"},
		{"en-US", "${lowerLocale}", "en-us"},
		{"en-US", "${upperLocale}", "EN-US"},

		{"de", "${base}", "de"},
		{"en-GB", "${base}", "en"},
		{"en-US", "${base}", "en"},
		{"en-US", "${lowerBase}", "en"},
		{"en-US", "${upperBase}", "EN"},

		{"de", "${region}", "DE"},
		{"en-GB", "${region}", "GB"},
		{"en-US", "${region}", "US"},
		{"en-US", "${lowerRegion}", "us"},
		{"en-US", "${upperRegion}", "US"},
	}
	for _, test := range testCases {
		actual := expandExportPath(test.name, test.path)
		if actual != test.expected {
			t.Errorf(
				"wrong expanded path: expected %s actual %s",
				strconv.Quote(test.expected),
				strconv.Quote(actual),
			)
		}
	}
}
