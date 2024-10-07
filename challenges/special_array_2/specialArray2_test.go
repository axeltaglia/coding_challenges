package main

import (
	"reflect"
	"testing"
)

func Test_isArraySpecial(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		nums    []int
		queries [][]int
		want    []bool
	}{
		{
			"Test 1: Single query, all valid",
			[]int{3, 4, 1, 2, 6},
			[][]int{{0, 4}},
			[]bool{false},
		},
		{
			"Test 2: Multiple queries, mixed results",
			[]int{4, 3, 1, 6},
			[][]int{{0, 2}, {2, 3}},
			[]bool{false, true},
		},
		{
			"Test 3: Query with no adjacent pairs",
			[]int{1},
			[][]int{{0, 0}},
			[]bool{true}, // Single element is trivially special
		},
		{
			"Test 4: All elements of the same parity",
			[]int{2, 4, 6, 8},
			[][]int{{0, 3}},
			[]bool{false},
		},
		{
			"Test 5: Alternating parity",
			[]int{1, 2, 3, 4, 5, 6},
			[][]int{{0, 5}},
			[]bool{true},
		},
		{
			"Test 6: Edge case - empty query",
			[]int{1, 2, 3},
			[][]int{},
			[]bool{},
		},
		{
			"Test 7: Subarray with mixed parity at edges",
			[]int{1, 2, 3, 5, 7},
			[][]int{{1, 3}},
			[]bool{true},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := isArraySpecial2(test.nums, test.queries); !reflect.DeepEqual(got, test.want) {
				t.Errorf("isArraySpecial(%v, %v) = %v, want %v", test.nums, test.queries, got, test.want)
			}
		})
	}
}
