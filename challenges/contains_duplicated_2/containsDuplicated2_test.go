package main

import "testing"

func Test_containsDuplicated2(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
		k    int
	}
	tests := []struct {
		name  string
		input input
		want  bool
	}{
		{
			"test1",
			input{
				[]int{1, 2, 3, 1},
				3,
			},
			true,
		},
		{
			"test2",
			input{
				[]int{1, 0, 1, 1},
				1,
			},
			true,
		},
		{
			"test3",
			input{
				[]int{1, 2, 3, 1, 2, 3},
				2,
			},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := containsNearbyDuplicate(test.input.nums, test.input.k); got != test.want {
				t.Errorf("test: %v, got: %v, want: %v", test.name, got, test.want)
			}
		})

	}
}
