package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		// Basic Cases
		{"Subsequence found", "abc", "ahbgdc", true},
		{"Subsequence not found", "axc", "ahbgdc", false},

		// Edge Cases
		{"Empty s, non-empty t", "", "abc", true},
		{"Non-empty s, empty t", "abc", "", false},
		{"Both s and t empty", "", "", true},

		// Special Cases
		{"s has more occurrences than t", "aaaa", "aa", false},
		{"Characters of s appear in t in order", "abc", "aabbcc", true},
		{"s partially matches t", "ace", "abcde", true},
		{"s does not match t", "aec", "abcde", false},

		// Larger Inputs
		{"Long subsequence", "abcdefgh", "aabbccddeeffgghh", true},
		{"Long t, short s", "ace", "abcdefghijklmnopqrstuvwxyz", true},
		{"No match in large input", "xyz", "abcdefghijklmnopqrstuvw", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isSubsequence(test.s, test.t)
			if result != test.expected {
				t.Errorf("For s = %q and t = %q, expected %v but got %v", test.s, test.t, test.expected, result)
			}
		})
	}
}
