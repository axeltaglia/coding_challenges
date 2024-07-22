package main

func main() {

}
func singleNumber(nums []int) int {
	acum := 0
	total := 0
	hash := make(map[int]bool, len(nums))

	for _, num := range nums {
		acum += num
		_, ok := hash[num]

		if !ok {
			total += num
			hash[num] = true
		}
	}

	return total*2 - acum
}
