package main

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Simple case",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "Trailing spaces",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Single word",
			input:    "joyboy",
			expected: 6,
		},
		{
			name:     "Multiple spaces in between",
			input:    "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "All spaces",
			input:    "         ",
			expected: 0, // Not valid per constraints (but added for robustness).
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0, // Not valid per constraints (but added for robustness).
		},
		{
			name:     "One word with trailing spaces",
			input:    "word    ",
			expected: 4,
		},
		{
			name:     "One word without spaces",
			input:    "word",
			expected: 4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lengthOfLastWord(test.input)
			if result != test.expected {
				t.Errorf("lengthOfLastWord(%q) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}
