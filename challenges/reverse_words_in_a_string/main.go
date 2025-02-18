package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	//s := []byte("  hola   mundo  ")
	//reverse(s)
	//fmt.Println(string(s))
}

//   odnum   aloh

func reverseWords(s string) string {
	inWord := false
	var currentWord []byte
	var arr []string

	for _, ch := range s {
		if !inWord && ch != ' ' {
			currentWord = append(currentWord, byte(ch))
			inWord = true
		} else if inWord && ch == ' ' {
			inWord = false
			//fmt.Println(string(currentWord))
			arr = append(arr, string(currentWord))
			currentWord = []byte{}
		} else if inWord {
			currentWord = append(currentWord, byte(ch))
		}
	}

	if len(currentWord) > 0 {
		arr = append(arr, string(currentWord))
	}

	var sb strings.Builder

	for i := len(arr) - 1; i >= 0; i-- {
		sb.WriteString(arr[i])
		if i > 0 {
			sb.WriteString(" ")
		}
	}

	fmt.Println(arr)
	return sb.String()
}

func reverse(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

// hola mundo
// odnu maloh
