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
		// add more cases here
		{
			input:    "HELLO wORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "helloWorld",
			expected: []string{"helloworld"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the len of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Number of elements is not matching. %d vs. expected %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("One or more elements are not matcing: %d is the first instance", i)
			}
		}
	}
}
