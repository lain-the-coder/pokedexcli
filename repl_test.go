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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// 1. Check if the number of words matches
		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch: got %d words, want %d words", len(actual), len(c.expected))
			continue // skip word check if lengths don't match to prevent index out of bounds
		}

		// 2. Check if each individual word matches
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word mismatch at index %d: got %q, want %q", i, word, expectedWord)
			}
		}
	}
}
