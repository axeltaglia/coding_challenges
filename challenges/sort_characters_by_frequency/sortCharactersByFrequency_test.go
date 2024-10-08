package main

import (
	"testing"
)

// Helper function to check if two strings have the same character frequency
func sameCharFrequency(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hash1 := make(map[rune]int)
	hash2 := make(map[rune]int)

	for _, c := range s1 {
		hash1[c]++
	}
	for _, c := range s2 {
		hash2[c]++
	}

	for k, v := range hash1 {
		if hash2[k] != v {
			return false
		}
	}
	return true
}

func TestFrequencySort(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"Test with multiple characters", "tree", "eert"}, // Can also be "eert" or "rtee"
		{"Test with single character", "aaa", "aaa"},
		{"Test with mixed frequencies", "cccaaa", "aaaccc"}, // "cccaaa" is also valid
		{"Test with numbers", "Aabb", "bbAa"},
		{"Test with special characters", "!a!a!b", "!!!aab"}, // or "aa!!!b"
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := frequencySort(tt.input)
			if !sameCharFrequency(got, tt.want) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
