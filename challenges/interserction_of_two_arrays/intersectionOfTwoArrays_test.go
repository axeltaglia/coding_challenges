package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "Test with one common element, different counts",
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			name:     "Test with multiple common elements",
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{9, 4},
		},
		{
			name:     "Test with no common elements",
			nums1:    []int{1, 3, 5, 7},
			nums2:    []int{2, 4, 6, 8},
			expected: []int{},
		},
		{
			name:     "Test with identical arrays",
			nums1:    []int{6, 6, 6},
			nums2:    []int{6, 6, 6},
			expected: []int{6},
		},
		{
			name:     "Test with completely different arrays",
			nums1:    []int{0, 1, 2},
			nums2:    []int{3, 4, 5},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := intersection(test.nums1, test.nums2)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("For %v and %v, expected %v, but got %v", test.nums1, test.nums2, test.expected, result)
			}
		})
	}
}
