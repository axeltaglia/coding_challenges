/*
3152. Special Array II
Solved
Medium
Topics
Companies
Hint
An array is considered special if every pair of its adjacent elements contains two numbers with different parity.

You are given an array of integer nums and a 2D integer matrix queries, where for queries[i] = [fromi, toi] your task is to check that
subarray
 nums[fromi..toi] is special or not.

Return an array of booleans answer such that answer[i] is true if nums[fromi..toi] is special.



Example 1:

Input: nums = [3,4,1,2,6], queries = [[0,4]]

Output: [false]

Explanation:

The subarray is [3,4,1,2,6]. 2 and 6 are both even.

Example 2:

Input: nums = [4,3,1,6], queries = [[0,2],[2,3]]

Output: [false,true]

Explanation:

The subarray is [4,3,1]. 3 and 1 are both odd. So the answer to this query is false.
The subarray is [1,6]. There is only one pair: (1,6) and it contains numbers with different parity. So the answer to this query is true.


Constraints:

1 <= nums.length <= 105
1 <= nums[i] <= 105
1 <= queries.length <= 105
queries[i].length == 2
0 <= queries[i][0] <= queries[i][1] <= nums.length - 1
*/

package main

import "fmt"

func main() {
	nums := []int{4, 3, 1, 6}
	queries := [][]int{{0, 2}, {2, 3}}
	fmt.Println(isArraySpecial2(nums, queries))
}

func isArraySpecial2(nums []int, queries [][]int) []bool {
	h := make(map[int]int)
	h[0] = 1
	for i := 1; i < len(nums); i++ {
		if differentParity(nums[i], nums[i-1]) {
			h[i] = h[i-1] + 1
		} else {
			h[i] = h[i-1]
		}
	}

	resul := make([]bool, 0)

	for _, query := range queries {
		from := query[0]
		to := query[1]

		if h[to]-h[from] == to-from {
			resul = append(resul, true)
		} else {
			resul = append(resul, false)
		}
	}

	return resul
}

func differentParity(n1 int, n2 int) bool {
	p1 := n1 % 2
	p2 := n2 % 2
	return p1 != p2
}
