package main

import "math"

func main() {
	maxProfit([]int{7, 1, 5, 3, 6, 4})
}
func maxProfit(prices []int) int {
	min := math.MaxInt
	maxProfit := 0
	for _, n := range prices {
		if n < min {
			min = n
		}
		if n-min > maxProfit {
			maxProfit = n - min
		}
	}

	return maxProfit
}
