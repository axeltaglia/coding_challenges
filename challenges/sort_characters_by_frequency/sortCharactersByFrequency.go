package main

import (
	"fmt"
	"sort"
	"strings"
)

func frequencySort(s string) string {
	// Create a hash map to store the frequency of each character
	hash := make(map[rune]int)
	for _, c := range s {
		hash[c]++
	}

	// Create a slice of runes (keys from the hash map)
	runes := make([]rune, 0, len(hash))
	for k := range hash {
		runes = append(runes, k)
	}

	// Sort the runes slice based on frequency in descending order
	sort.Slice(runes, func(i, j int) bool {
		return hash[runes[i]] > hash[runes[j]]
	})

	// Build the output string
	var output strings.Builder
	for _, r := range runes {
		output.WriteString(strings.Repeat(string(r), hash[r]))
	}

	return output.String()
}

func main() {
	s := "tree"
	fmt.Println(frequencySort(s)) // Output could be: "eert" or "rtee"
}
