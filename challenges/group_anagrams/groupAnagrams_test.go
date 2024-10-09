package main

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected [][]string
	}{
		{
			name:     "Test with multiple anagram groups",
			input:    []string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"},
			expected: [][]string{{"cab"}, {"tin"}, {"pew"}, {"duh"}, {"may"}, {"ill"}, {"buy"}, {"bar"}, {"max"}, {"doc"}},
		},
		{
			name:     "Test with no anagrams",
			input:    []string{"abc", "def", "ghi"},
			expected: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name:     "Test with repeated anagrams",
			input:    []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := groupAnagrams(tt.input)
			if !equalIgnoringOrder(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// Helper function to compare without considering the order
func equalIgnoringOrder(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, groupA := range a {
		found := false
		for _, groupB := range b {
			if reflect.DeepEqual(groupA, groupB) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
