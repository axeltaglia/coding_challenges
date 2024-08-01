/*
SPECIAL ARRAY

An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integers nums. Return true if nums is a special array, otherwise, return false.



Example 1:

Input: nums = [1]

Output: true

Explanation:

There is only one element. So the answer is true.

Example 2:

Input: nums = [2,1,4]

Output: true

Explanation:

There is only two pairs: (2,1) and (1,4), and both of them contain numbers with different parity. So the answer is true.

Example 3:

Input: nums = [4,3,1,6]

Output: false

Explanation:

nums[1] and nums[2] are both odd. So the answer is false.



Constraints:

1 <= nums.length <= 100
1 <= nums[i] <= 100
*/

package main

func isArraySpecial(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if sameParity(nums[i], nums[i+1]) {
			return false
		}
	}
	return true
}

func sameParity(n1 int, n2 int) bool {
	return n1%2 == n2%2
}
