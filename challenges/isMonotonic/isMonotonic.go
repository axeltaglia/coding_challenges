package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 3, 5, 6}))
}
func isMonotonic(nums []int) bool {
	isMonotonicInc := true
	isMonotonicDec := true

	for i := 1; i < len(nums); i++ {
		if isMonotonicInc && nums[i] < nums[i-1] {
			isMonotonicInc = false
		}
		if isMonotonicDec && nums[i] > nums[i-1] {
			isMonotonicDec = false
		}
	}

	return isMonotonicInc || isMonotonicDec
}
