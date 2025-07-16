package main

import (
	"testing"
)



func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  ash   ketchum  ",
			expected: []string{"ash", "ketchum"},
		},
		// add more test cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d words, expected %d words", c.input, len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				t.Errorf("cleanInput(%q) word %d = %q, expected %q", c.input, i, word, expectedWord)
			}
		}
	}


}