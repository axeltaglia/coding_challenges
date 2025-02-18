package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return []int{0}
	}

	leftProduct := make([]int, len(nums))
	leftProduct[0] = 1
	acum := nums[0]
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = acum
		acum *= nums[i]
	}

	rightProduct := make([]int, len(nums))
	rightProduct[len(rightProduct)-1] = 1
	acum = nums[len(rightProduct)-1]
	for i := len(rightProduct) - 2; i >= 0; i-- {
		rightProduct[i] = acum
		acum *= nums[i]
	}

	for i := 0; i < len(leftProduct); i++ {
		leftProduct[i] *= rightProduct[i]
	}

	return leftProduct
}

// 2 3 2 1
// 1 2 6 12
// | | | | _ _ 1X2X3 = 6
// | | | _ _ _ 2x3 = 6     acum=nums[i]*acum=2*6=12
// | | _ _ _ _ 2           acum=nums[i]*acum= 2 * 3 = 6
// | _ _ _ _ _ 1 (fixed)   acum=2

// 2 3 2 1
// 1

// ini:
// leftProd[0] = 1
// acum = nums[0]

// for i:=1;...
// leftProd[i] = acum
// acum *= nums[i]
