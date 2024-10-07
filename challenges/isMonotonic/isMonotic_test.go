package main

import (
	"testing"
)

func Test_isMonotonic(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"Monotonic increasing", []int{1, 2, 2, 3}, true},
		{"Monotonic decreasing", []int{6, 5, 4, 4}, true},
		{"Not monotonic", []int{1, 3, 2}, false},
		{"All elements the same", []int{3, 3, 3, 3}, true},
		{"Single element", []int{10}, true},
		{"Monotonic strictly increasing", []int{1, 2, 3, 4, 5}, true},
		{"Monotonic strictly decreasing", []int{5, 4, 3, 2, 1}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMonotonic(tt.nums)
			if got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
