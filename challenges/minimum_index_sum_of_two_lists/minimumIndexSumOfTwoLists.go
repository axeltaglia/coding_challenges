package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}
func findRestaurant(list1 []string, list2 []string) []string {
	indexSums := make(map[string]int, 0)
	for pos, hotel := range list1 {
		indexSums[hotel] = pos
	}

	output := make([]string, 0)
	min := math.MaxInt32
	for i, hotel := range list2 {
		if _, exists := indexSums[hotel]; exists {
			indexSum := i + indexSums[hotel]
			if indexSum < min {
				min = indexSum
				output = []string{hotel}
			} else if indexSum == min {
				output = append(output, hotel)
			}
		}
	}

	return output
}
