package main

import "fmt"

func main() {
	fmt.Println(countHillValley([]int{1, 2, 3, 2, 3, 2}))
}

func countHillValley(nums []int) int {
	hills := 0
	valleys := 0
	lastTransition := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] { // equals
			continue
		} else if nums[i] > nums[i-1] { // increasing
			if lastTransition == -1 {
				valleys++
			}
			lastTransition = 1
		} else if nums[i] < nums[i-1] { // decreasing
			if lastTransition == 1 {
				hills++
			}
			lastTransition = -1
		}
	}

	return hills + valleys
}
