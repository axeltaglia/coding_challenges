package main

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Simple case III",
			input:    "III",
			expected: 3,
		},
		{
			name:     "Simple case LVIII",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "Complex case MCMXCIV",
			input:    "MCMXCIV",
			expected: 1994,
		},
		{
			name:     "Single numeral I",
			input:    "I",
			expected: 1,
		},
		{
			name:     "Single numeral M",
			input:    "M",
			expected: 1000,
		},
		{
			name:     "Subtractive notation IV",
			input:    "IV",
			expected: 4,
		},
		{
			name:     "Subtractive notation IX",
			input:    "IX",
			expected: 9,
		},
		{
			name:     "Subtractive notation XL",
			input:    "XL",
			expected: 40,
		},
		{
			name:     "Subtractive notation XC",
			input:    "XC",
			expected: 90,
		},
		{
			name:     "Subtractive notation CD",
			input:    "CD",
			expected: 400,
		},
		{
			name:     "Subtractive notation CM",
			input:    "CM",
			expected: 900,
		},
		{
			name:     "Mixed notation LXXXIII",
			input:    "LXXXIII",
			expected: 83,
		},
		{
			name:     "Mixed notation MMXXIV",
			input:    "MMXXIV",
			expected: 2024,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := romanToInt(test.input)
			if result != test.expected {
				t.Errorf("romanToInt(%s) = %d; expected %d", test.input, result, test.expected)
			}
		})
	}
}
