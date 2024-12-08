package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Common prefix exists",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "No common prefix",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "Single string",
			input:    []string{"single"},
			expected: "single",
		},
		{
			name:     "Empty array",
			input:    []string{},
			expected: "",
		},
		{
			name:     "All strings are the same",
			input:    []string{"test", "test", "test"},
			expected: "test",
		},
		{
			name:     "One empty string",
			input:    []string{"", "abc", "abd"},
			expected: "",
		},
		{
			name:     "Strings with different lengths",
			input:    []string{"abcd", "abc", "ab"},
			expected: "ab",
		},
		{
			name:     "Common prefix with one character",
			input:    []string{"a", "ab", "abc"},
			expected: "a",
		},
		{
			name:     "No strings with common characters",
			input:    []string{"x", "y", "z"},
			expected: "",
		},
		{
			name:     "Large input with common prefix",
			input:    []string{"prefix_common", "prefix_test", "prefix_longer"},
			expected: "prefix_",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := longestCommonPrefix(test.input)
			if result != test.expected {
				t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
			}
		})
	}
}
