package main

import (
	"testing"
)

func Test_countHillValley(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		// Example cases
		{"Test with hills and valleys", []int{2, 4, 1, 1, 6, 5}, 3}, // 4->1 (valley), 1->6 (hill), 6->5 (valley)
		{"Test with hills only", []int{1, 2, 3, 2, 3, 2}, 3},        // 2->3->2 (hill), 3->2->3 (valley), 2->3->2 (hill)
		{"Test with no hills or valleys", []int{1, 1, 1, 1}, 0},     // No hills or valleys
		{"Test with increasing sequence", []int{1, 2, 3, 4}, 0},     // No hills or valleys in a fully increasing sequence
		{"Test with decreasing sequence", []int{4, 3, 2, 1}, 0},     // No hills or valleys in a fully decreasing sequence
		{"Test with flat sequence", []int{2, 2, 2}, 0},              // No hills or valleys
		{"Test with single valley", []int{5, 1, 5}, 1},              // 1 valley at 5->1->5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countHillValley(tt.nums); got != tt.want {
				t.Errorf("countHillValley() = %v, want %v", got, tt.want)
			}
		})
	}
}
