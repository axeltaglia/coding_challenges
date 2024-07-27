package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Parallel()

	type input struct {
		nums   []int
		target int
	}

	tests := []struct {
		name  string
		input input
		want  []int
	}{
		{
			"Test 1",
			input{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			"Test 2",
			input{[]int{3, 2, 4}, 6},
			[]int{1, 2},
		},
		{
			"Test 3",
			input{[]int{3, 3}, 6},
			[]int{0, 1},
		},
		{
			"Test 4",
			input{[]int{1, 2, 3, 4, 5}, 10},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := twoSum(test.input.nums, test.input.target); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
