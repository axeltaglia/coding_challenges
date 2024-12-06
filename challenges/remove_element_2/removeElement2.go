package main

func main() {
	removeDuplicates([]int{})
}
func removeDuplicates(nums []int) int {
	c := 0
	r := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i-1] == nums[i] {
			r++
		} else {
			r = 1
		}

		if r <= 2 {
			nums[c] = nums[i]
			c++
		}
	}

	return c
}
