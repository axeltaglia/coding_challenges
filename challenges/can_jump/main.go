/*
55. Jump Game
Attempted
Medium
Topics
Companies
You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

Return true if you can reach the last index, or false otherwise.



Example 1:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


Constraints:

1 <= nums.length <= 104
0 <= nums[i] <= 105

*/

package main

import "fmt"

func main() {
	//fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{1, 2, 3}))
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}
	maxReach := 0
	for i, n := range nums {
		if i > maxReach {
			return false
		}
		if i+n > maxReach {
			maxReach = i + n
		}
		if i+n >= len(nums)-1 {
			return true
		}

	}
	return false
}

func canJump2(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}
	c := 0
	for c < len(nums) {
		if c == len(nums)-1 {
			return true
		}
		if nums[c] == 0 {
			return false
		}

		c = c + nums[c]
	}

	return false
}

//   .
// 2,3,1,1,4
