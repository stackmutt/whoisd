package storage

import (
	"testing"
)

func TestCustomJoin(t *testing.T) {

	type testData struct {
		format   string
		value    []string
		expected string
	}

	var tests = []testData{
		{"{string}.{string}{string}", []string{"+31", "1025", "1112"}, "+31.10251112"},
		{"{string}.{string}{string}", []string{"", "1025", "1112"}, "10251112"},
		{"{date}", []string{"2014-06-01 12:20:16"}, "2014-06-01T12:20:16Z"},
		{"Last update of WHOIS database: {date}",
			[]string{"2014-06-01 12:20:16"},
			"Last update of WHOIS database: 2014-06-01T12:20:16Z",
		},
	}

	for _, data := range tests {
		result := customJoin(data.format, data.value)
		if result != data.expected {
			t.Error("Expected ", data.expected, ", got", result)
		}
	}
}
