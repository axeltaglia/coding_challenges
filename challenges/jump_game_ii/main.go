package main

import "fmt"

func main() {
	fmt.Println(jump([]int{2, 1}))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	jumps := 0
	maxReach := 0
	currentEnd := 0

	for i := 0; i < len(nums); i++ {
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}

		if i == currentEnd {
			currentEnd = maxReach
			jumps++
			if maxReach >= len(nums)-1 {
				return jumps
			}
		}
	}

	return -1
}

//   .
// 2,3,1,1,4

// mR = 4
// cE = 2
