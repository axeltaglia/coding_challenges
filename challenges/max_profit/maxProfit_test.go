package main

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"Normal case", []int{7, 1, 5, 3, 6, 4}, 5},
		{"Decreasing prices", []int{7, 6, 4, 3, 1}, 0},
		{"Single price", []int{5}, 0},
		{"All same prices", []int{3, 3, 3, 3}, 0},
		{"Increasing prices", []int{1, 2, 3, 4, 5}, 4},
		{"Mixed prices", []int{2, 4, 1, 5, 3, 6}, 5},
		{"Large profit opportunity", []int{1, 10}, 9},
		{"Large input, no profit", []int{10000, 9000, 8000, 7000, 6000}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
