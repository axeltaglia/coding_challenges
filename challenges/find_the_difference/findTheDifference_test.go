package main

import (
	"reflect"
	"testing"
)

func Test_findTheDifference(t *testing.T) {
	t.Parallel()

	type input struct {
		s string
		t string
	}

	tests := []struct {
		name  string
		input input
		want  byte
	}{
		{
			"Test 1",
			input{"abcd", "abcde"},
			'e',
		},
		{
			"Test 2",
			input{"", "y"},
			'y',
		},
		{
			"Test 3",
			input{"aabbcc", "abccbaa"},
			'a',
		},
		{
			"Test 4",
			input{"xyz", "yxzz"},
			'z',
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := findTheDifference(test.input.s, test.input.t); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
