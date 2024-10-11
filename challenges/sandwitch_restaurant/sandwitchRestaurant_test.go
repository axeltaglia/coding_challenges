package main

import (
	"testing"
)

func TestSandwichRestaurant(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected float64
	}{
		{
			name:     "Test with simple case",
			input:    [][]int{{0, 3}, {1, 9}, {2, 6}},
			expected: 10.0,
		},
		{
			name:     "Test with no extra waiting time",
			input:    [][]int{{0, 2}, {2, 4}, {6, 3}},
			expected: 3.0,
		},
		{
			name:     "Test with one customer coming late",
			input:    [][]int{{0, 5}, {1, 6}, {10, 2}},
			expected: 6,
		},
		{
			name:     "Test with all customers arriving at the same time",
			input:    [][]int{{0, 3}, {0, 2}, {0, 6}},
			expected: 6.33,
		},
		{
			name:     "Test with customers arriving far apart",
			input:    [][]int{{0, 2}, {10, 5}, {15, 3}},
			expected: 3.33,
		},
		{
			name:     "Test with short sandwich times",
			input:    [][]int{{0, 1}, {0, 1}, {0, 1}},
			expected: 2.0,
		},
		{
			name:     "Test with one customer",
			input:    [][]int{{5, 3}},
			expected: 3.0,
		},
		{
			name:     "Test with random arrival and sandwich times",
			input:    [][]int{{1, 4}, {3, 3}, {5, 2}},
			expected: 4.67,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sandwichRestaurant(tt.input)
			if !almostEqual(result, tt.expected) {
				t.Errorf("expected %.2f, got %.2f", tt.expected, result)
			}
		})
	}
}

// Helper function to compare floating-point results within a small tolerance
func almostEqual(a, b float64) bool {
	const tolerance = 0.01
	return (a-b) < tolerance && (b-a) < tolerance
}
