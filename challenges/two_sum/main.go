package main

import "fmt"

func main() {
	result := twoSum([]int{4, 3, 2, 5, 6, 7}, 9)
	fmt.Print(result)
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for currentPosition, n := range nums {
		complement := target - n
		if complementPosition, ok := hash[complement]; ok {
			return []int{currentPosition, complementPosition}
		}

		hash[n] = currentPosition
	}

	return nil
}
