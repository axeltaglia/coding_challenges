package main

import "testing"

func Test_missingNumber(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
	}
	tests := []struct {
		name  string
		input input
		want  int
	}{
		{
			"test1",
			input{
				[]int{2, 2, 1},
			},
			1,
		},
		{
			"test2",
			input{
				[]int{4, 1, 2, 1, 2},
			},
			4,
		},
		{
			"test3",
			input{
				[]int{1},
			},
			1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := singleNumber(test.input.nums); got != test.want {
				t.Errorf("test: %v, got: %v, want: %v", test.name, got, test.want)
			}
		})
	}
}
