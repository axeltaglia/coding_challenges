package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindRestaurant(t *testing.T) {
	tests := []struct {
		name     string
		list1    []string
		list2    []string
		expected []string
	}{
		{
			name:     "Test with one common restaurant",
			list1:    []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:    []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
			expected: []string{"Shogun"},
		},
		{
			name:     "Test with two common restaurants, minimal sum of indexes",
			list1:    []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:    []string{"KFC", "Shogun", "Burger King"},
			expected: []string{"Shogun"},
		},
		{
			name:     "Test with two common restaurants, different index sums",
			list1:    []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:    []string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
			expected: []string{"KFC", "Burger King", "Tapioca Express", "Shogun"},
		},
		{
			name:     "Test with no common restaurants",
			list1:    []string{"Shogun", "Tapioca Express", "Burger King", "KFC"},
			list2:    []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse"},
			expected: []string{},
		},
		{
			name:     "Test with identical lists",
			list1:    []string{"Shogun", "KFC", "Tapioca Express"},
			list2:    []string{"Shogun", "KFC", "Tapioca Express"},
			expected: []string{"Shogun"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findRestaurant(tt.list1, tt.list2)

			// Sort both slices to ignore order
			sort.Strings(result)
			sort.Strings(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, expected %v", result, tt.expected)
			}
		})
	}
}
