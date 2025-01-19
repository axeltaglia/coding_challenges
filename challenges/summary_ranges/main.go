package main

import "fmt"

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	from := 0
	for i := 0; i < len(nums); i++ {
		if i+1 == len(nums) || i+1 < len(nums) && nums[i+1]-nums[i] > 1 {
			if i == from {
				result = append(result, fmt.Sprintf("%d", nums[from]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[from], nums[i]))
			}
			from = i + 1
		}
	}

	return result
}

//     .
// 0,1,2,4,5,7
// .

func summaryRanges2(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	result := make([]string, 0)
	i := 1
	from := 0
	for i <= len(nums) {
		if i == len(nums) || nums[i]-nums[i-1] > 1 {
			if i == from+1 {
				result = append(result, fmt.Sprintf("%d", nums[from]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[from], nums[i-1]))
			}
			from = i
		}
		i++
	}

	return result
}
