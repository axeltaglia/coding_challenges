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
				[]int{3, 0, 1},
			},
			2,
		},
		{
			"test2",
			input{
				[]int{0, 1},
			},
			2,
		},
		{
			"test3",
			input{
				[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			},
			8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := missingNumber(test.input.nums); got != test.want {
				t.Errorf("test: %v, got: %v, want: %v", test.name, got, test.want)
			}
		})

	}
}
