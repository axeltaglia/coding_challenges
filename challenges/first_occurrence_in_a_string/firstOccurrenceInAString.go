package main

import "fmt"

func main() {
	fmt.Println(strStr("mississipi", "issip"))
}
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	i := 0
	l := 0
	lsp := make([]int, len(needle))
	computeLsp(needle, lsp)

	for i < len(haystack) {
		if haystack[i] == needle[l] {
			l++
			i++
			if l == len(needle) {
				return i - len(needle)
			}
		} else {
			if l == 0 {
				i++
			} else {
				l = lsp[l-1]
			}
		}
	}

	return -1
}

func computeLsp(str string, lsp []int) {
	length := 0
	i := 1
	for i < len(str) {
		if str[i] == str[length] {
			length++
			lsp[i] = length
			i++
		} else {
			if length > 0 {
				length = lsp[length-1]
			} else {
				lsp[i] = 0
				i++
			}
		}
	}
}

// aabaabaaa
// 010123452
// 012345678
