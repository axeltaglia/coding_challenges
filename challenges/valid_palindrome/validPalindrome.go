package main

import "strings"

func isPalindrome(s string) bool {
	str := ""
	strLower := strings.ToLower(s)
	for _, ch := range strLower {
		if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
			str += string(ch)
		}
	}

	if str == "" {
		return true
	}

	middle := len(str) / 2

	j := len(str) - 1
	for i := 0; i <= middle; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}

	return true
}
