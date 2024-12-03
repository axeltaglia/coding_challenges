package main

import "testing"

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"Single element", []int{1}, 1},
		{"All the same", []int{2, 2, 2, 2}, 2},
		{"Majority in middle", []int{1, 2, 3, 2, 2, 2}, 2},
		{"Majority at end", []int{3, 3, 4, 3, 3}, 3},
		{"Negative numbers", []int{-1, -1, -1, 2, 2, -1}, -1},
		{"Large majority", []int{10, 10, 10, 5, 10, 10, 10}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement(tt.nums)
			if got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
