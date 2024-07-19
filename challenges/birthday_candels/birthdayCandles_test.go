package main

import (
	"reflect"
	"testing"
)

func Test_birthdayCandles(t *testing.T) {
	t.Parallel()

	type input struct {
		candles []int
	}

	tests := []struct {
		name  string
		input input
		want  int
	}{
		{
			"Test 1",
			input{[]int{2, 3, 1, 3}},
			2,
		},
		{
			"Test 2",
			input{[]int{1, 1, 1, 1}},
			4,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := birthdayCakeCandles(test.input.candles); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
