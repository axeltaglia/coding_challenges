/*
FIND THE DIFFERENCE

You are given two strings s and t.

String t is generated by random shuffling string s and then add one more letter at a random position.

Return the letter that was added to t.



Example 1:

Input: s = "abcd", t = "abcde"
Output: "e"
Explanation: 'e' is the letter that was added.
Example 2:

Input: s = "", t = "y"
Output: "y"


Constraints:

0 <= s.length <= 1000
t.length == s.length + 1
s and t consist of lowercase English letters.

*/

package main

import "fmt"

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(string(findTheDifference(s, t)))
}

func findTheDifference(s string, t string) byte {
	var counter int32

	for _, ch := range s {
		counter += ch
	}

	for _, ch := range t {
		counter -= ch
	}

	return byte(counter * -1)
}

func findTheDifference2(s string, t string) byte {
	hash := make(map[int32]int32)

	for _, ch := range s {
		hash[ch]++
	}

	for _, ch := range t {
		if _, ok := hash[ch]; !ok {
			return byte(ch)
		}

		hash[ch]--
		val := hash[ch]

		if val < 0 {
			return byte(ch)
		}
	}

	return byte(0)
}

// a, a, b, c
// a, b, a, c, a

// a: -1 --->
// b: 0
// c: 0

// a=1, b=2, c=3, d=4
// a, a, b, c
// 1+1+2+3=7
// a, b, a, a, c
// 6  4  3  2  -1
