/*
s Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/

package main

import "fmt"

func main() {
	fmt.Println(unique("asdkvmdb"))
}

func unique(s string) bool {
	dic := [128]bool{}
	for _, ch := range s {
		if dic[ch] {
			return false
		}
		dic[ch] = true
	}
	return true
}
