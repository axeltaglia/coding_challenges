/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the
other.
*/

package main

import "fmt"

func main() {
	fmt.Println(checkPermutation("sad", "dat"))
}

func checkPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	dic := [128]int{}

	for _, ch := range s1 {
		dic[ch]++
	}

	for _, ch := range s2 {
		dic[ch]--
		if dic[ch] < 0 {
			return false
		}
	}

	return true
}
