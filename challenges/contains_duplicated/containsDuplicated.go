/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

package main

func main() {
	println(containsDuplicate([]int{2, 1, 3, 5}))
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]bool)
	for _, n := range nums {
		if _, ok := hash[n]; ok {
			return true
		}
		hash[n] = true
	}

	return false
}
