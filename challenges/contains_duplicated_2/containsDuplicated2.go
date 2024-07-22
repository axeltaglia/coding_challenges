package main

func main() {

}

func containsNearbyDuplicate(nums []int, k int) bool {
	hash := make(map[int]int)

	for i, num := range nums {
		if j, ok := hash[num]; ok {
			if i-j <= k {
				return true
			}
		} else {
			hash[num] = i
		}
	}

	return false
}
