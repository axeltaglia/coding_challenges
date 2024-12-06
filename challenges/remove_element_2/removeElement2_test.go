package main

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expected  []int
		expectedK int
	}{
		{
			name:      "No duplicates",
			input:     []int{1, 2, 3},
			expected:  []int{1, 2, 3},
			expectedK: 3,
		},
		{
			name:      "Duplicates within allowed limit",
			input:     []int{1, 1, 2, 2, 3, 3},
			expected:  []int{1, 1, 2, 2, 3, 3},
			expectedK: 6,
		},
		{
			name:      "Duplicates exceeding allowed limit",
			input:     []int{1, 1, 1, 2, 2, 3},
			expected:  []int{1, 1, 2, 2, 3},
			expectedK: 5,
		},
		{
			name:      "Single element",
			input:     []int{1},
			expected:  []int{1},
			expectedK: 1,
		},
		{
			name:      "All identical elements",
			input:     []int{1, 1, 1, 1},
			expected:  []int{1, 1},
			expectedK: 2,
		},
		{
			name:      "Mixed duplicates",
			input:     []int{0, 0, 1, 1, 1, 2, 3, 3, 3},
			expected:  []int{0, 0, 1, 1, 2, 3, 3},
			expectedK: 7,
		},
		{
			name:      "Empty array",
			input:     []int{},
			expected:  []int{},
			expectedK: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.input...) // Copy input to avoid mutation
			k := removeDuplicates(nums)

			if k != tt.expectedK {
				t.Errorf("removeDuplicates() returned k = %d, expected %d", k, tt.expectedK)
			}

			if !reflect.DeepEqual(normalizeSlice(nums[:k]), normalizeSlice(tt.expected)) {
				t.Errorf("removeDuplicates() modified nums = %v, expected %v", nums[:k], tt.expected)
			}
		})
	}
}

func normalizeSlice(slice []int) []int {
	if slice == nil {
		return []int{}
	}
	return slice
}
