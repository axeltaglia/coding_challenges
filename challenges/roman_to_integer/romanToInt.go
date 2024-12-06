package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	sum := 0

	hash := make(map[rune]int)
	hash['I'] = 1
	hash['V'] = 5
	hash['X'] = 10
	hash['L'] = 50
	hash['C'] = 100
	hash['D'] = 500
	hash['M'] = 1000

	for i, ch := range s {
		if i < len(s)-1 && hash[ch] < hash[rune(s[i+1])] {
			sum -= hash[ch]
		} else {
			sum += hash[ch]
		}
	}

	return sum
}
