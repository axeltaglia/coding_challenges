package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}
func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]int) // Store the last index of each character
	start, maxLen := 0, 0      // `start` is the beginning of the sliding window

	for i, ch := range s {
		if lastIndex, ok := hash[ch]; ok && lastIndex >= start {
			// If character is repeated, move start to the right of the last occurrence
			start = lastIndex + 1
		}
		// Calculate the current window length and update maxLen if needed
		currentLen := i - start + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
		// Update the last seen index of the current character
		hash[ch] = i
	}

	return maxLen
}
