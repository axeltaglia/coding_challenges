package main

import (
	"reflect"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]string
		expected [][]string
	}{
		{
			name: "Test with multiple accounts and common emails",
			input: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
				{"Mary", "mary@mail.com"},
			},
			expected: [][]string{
				{"John", "johnsmith@mail.com", "john00@mail.com", "john_newyork@mail.com"},
				{"John", "johnnybravo@mail.com"},
				{"Mary", "mary@mail.com"},
			},
		},
		{
			name: "Test with no common emails",
			input: [][]string{
				{"Alice", "alice@mail.com"},
				{"Bob", "bob@mail.com"},
			},
			expected: [][]string{
				{"Alice", "alice@mail.com"},
				{"Bob", "bob@mail.com"},
			},
		},
		{
			name: "Test with identical emails",
			input: [][]string{
				{"Alice", "alice@mail.com", "alice@mail.com"},
			},
			expected: [][]string{
				{"Alice", "alice@mail.com"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := accountsMerge(tt.input)
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
