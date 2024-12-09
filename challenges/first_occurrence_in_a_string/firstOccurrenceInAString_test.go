package main

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
		{"abc", "", 0},
		{"", "a", -1},
		{"mississippi", "issip", 4},
		{"mississippi", "issipi", -1},
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"aaaaa", "aaaaa", 0},
		{"aabaaabaaac", "aabaaac", 4},
	}

	for _, tt := range tests {
		t.Run(tt.haystack+"_"+tt.needle, func(t *testing.T) {
			result := strStr(tt.haystack, tt.needle)
			if result != tt.expected {
				t.Errorf("strStr(%q, %q) = %d; expected %d", tt.haystack, tt.needle, result, tt.expected)
			}
		})
	}
}
