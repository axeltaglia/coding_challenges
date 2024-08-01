package main

import (
	"reflect"
	"testing"
)

func Test_isArraySpecial(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
	}

	tests := []struct {
		name  string
		input input
		want  bool
	}{
		{
			"Test 1",
			input{[]int{1, 2, 3, 4, 5}},
			true,
		},
		{
			"Test 2",
			input{[]int{1, 1, 2, 2}},
			false,
		},
		{
			"Test 3",
			input{[]int{1, 2, 3, 4, 6}},
			false,
		},
		{
			"Test 4",
			input{[]int{2, 4, 6, 8, 10}},
			false,
		},
		{
			"Test 5",
			input{[]int{1, 3, 5, 7, 9}},
			false,
		},
		{
			"Test 6",
			input{[]int{1, 2, 4, 5, 7, 8}},
			false,
		},
		{
			"Test 7",
			input{[]int{1}},
			true,
		},
		{
			"Test 8",
			input{[]int{}},
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := isArraySpecial(test.input.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
