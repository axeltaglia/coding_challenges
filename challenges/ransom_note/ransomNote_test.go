package main

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		// Basic cases
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},

		// Edge cases
		{"", "", true},         // Empty ransom note and magazine
		{"", "anything", true}, // Empty ransom note
		{"some", "", false},    // Empty magazine
		{"a", "a", true},       // Single character match
		{"aaaa", "aaaa", true}, // Exact match
		{"aaaa", "aaa", false}, // Not enough characters

		// Complex cases
		{"abc", "abccba", true},    // Sufficient characters in magazine
		{"abc", "abcdabc", true},   // Multiple occurrences in magazine
		{"abcabc", "aabbcc", true}, // Ransom note has repeated characters
		{"abcabc", "aabbc", false}, // Insufficient characters for repeated ransom note

	}

	for _, test := range tests {
		t.Run(test.ransomNote+"_"+test.magazine, func(t *testing.T) {
			result := canConstruct(test.ransomNote, test.magazine)
			if result != test.expected {
				t.Errorf("For ransomNote '%s' and magazine '%s', expected %v but got %v",
					test.ransomNote, test.magazine, test.expected, result)
			}
		})
	}
}
