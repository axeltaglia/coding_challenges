/*

Code


Testcase
Testcase
Test Result
349. Intersection of Two Arrays
Solved
Easy
Topics
Companies
Given two integer arrays nums1 and nums2, return an array of their
intersection
. Each element in the result must be unique and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4]
Explanation: [4,9] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000
*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	result := intersection(nums1, nums2)
	fmt.Println(result) // Output: [2]

	nums3 := []int{4, 9, 5}
	nums4 := []int{9, 4, 9, 8, 4}
	result2 := intersection(nums3, nums4)
	fmt.Println(result2) // Output: [9, 4] or [4, 9]
}

func intersection(nums1 []int, nums2 []int) []int {
	output := make([]int, 0)
	hashNums1 := make(map[int]bool)
	for _, n := range nums1 {
		hashNums1[n] = true
	}

	for _, n := range nums2 {
		if hashNums1[n] {
			output = append(output, n)
			delete(hashNums1, n)
		}
	}
	return output
}
