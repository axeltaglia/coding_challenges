package main

import "fmt"

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}
func findRestaurant(list1 []string, list2 []string) []string {
	hotelFreqs := make(map[string]int, 0)
	acumPositions := make(map[string]int, 0)
	for pos, hotel := range list1 {
		hotelFreqs[hotel]++
		acumPositions[hotel] += pos
	}

	for pos, hotel := range list2 {
		hotelFreqs[hotel]++
		acumPositions[hotel] += pos
	}

	min := 100000000
	for hotel, freq := range hotelFreqs {
		if freq == 2 {
			if acumPositions[hotel] < min {
				min = acumPositions[hotel]
			}
		} else {
			acumPositions[hotel] = -1
		}
	}

	output := make([]string, 0)

	for hotel, acumPosition := range acumPositions {
		if acumPosition == min {
			output = append(output, hotel)
		}
	}

	return output
}
