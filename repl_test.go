package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Charmander Bulbasaur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "one, two, three",
			expected: []string{"one,", "two,", "three"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			// error and continue here
			t.Errorf("expected len: %v, got: %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("expected: %v, got: %v", expectedWord, word)
			}
		}
	}
}
