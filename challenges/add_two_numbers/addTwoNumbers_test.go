package main

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Parallel()

	type input struct {
		l1     *ListNode
		l2     *ListNode
		carrie int
	}

	tests := []struct {
		name  string
		input input
		want  *ListNode
	}{
		{
			"Test 1",
			input{
				&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
				&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
				0,
			},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			"Test 2",
			input{
				&ListNode{0, nil},
				&ListNode{0, nil},
				0,
			},
			&ListNode{0, nil},
		},
		{
			"Test 3",
			input{
				&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
				&ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
				0,
			},
			&ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
		{
			"Test 4",
			input{
				nil,
				nil,
				0,
			},
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			if got := addTwoNumbers(test.input.l1, test.input.l2, test.input.carrie); !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
