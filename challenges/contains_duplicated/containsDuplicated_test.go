package main

import (
	"reflect"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
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
			input{[]int{1, 2, 3, 1}},
			true,
		},
		{
			"Test 2",
			input{[]int{1, 2, 3, 4}},
			false,
		},
		{
			"Test 3",
			input{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			true,
		},
		{
			"Test 4",
			input{[]int{}},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := containsDuplicate(test.input.nums); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
