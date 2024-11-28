package main

import (
	"reflect"
	"testing"
)

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		val      int
		wantNums []int
		wantLen  int
	}{
		{"Remove from middle", []int{3, 2, 2, 3}, 3, []int{2, 2}, 2},
		{"Remove all elements", []int{3, 3, 3}, 3, []int{}, 0},
		{"Remove none", []int{1, 2, 3}, 4, []int{1, 2, 3}, 3},
		{"Mixed removal", []int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 1, 3, 0, 4}, 5},
		{"Single match", []int{1}, 1, []int{}, 0},
		{"Single no match", []int{1}, 2, []int{1}, 1},
		{"Empty array", []int{}, 1, []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := append([]int(nil), tt.nums...) // Create a copy to avoid modifying the original test case
			gotLen := removeElement(numsCopy, tt.val)
			gotNums := numsCopy[:gotLen]

			if gotLen != tt.wantLen {
				t.Errorf("removeElement() length = %d, want %d", gotLen, tt.wantLen)
			}
			if !reflect.DeepEqual(gotNums, tt.wantNums) {
				t.Errorf("removeElement() nums = %v, want %v", gotNums, tt.wantNums)
			}
		})
	}
}
