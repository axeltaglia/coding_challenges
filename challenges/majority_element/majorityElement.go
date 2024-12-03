package main

func main() {
	majorityElement([]int{8, 8, 7, 7, 7})
}

func majorityElement(nums []int) int {
	var candidate int
	var count int

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}

		if n == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
