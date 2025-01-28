package main

import "fmt"

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

func wordPattern(pattern string, s string) bool {
	hash := make(map[rune]string)
	hash2 := make(map[string]rune)

	var c int
	start := 0
	for i := 0; i <= len(s); i++ {

		if i == len(s) || rune(s[i]) == ' ' {
			if c == len(pattern) {
				return false
			}
			word := s[start:i]
			currentPatLetter := rune(pattern[c])

			expectedWord, ok := hash[currentPatLetter]
			if ok && word != expectedWord {
				return false
			}

			expectedPatLetter, ok := hash2[word]
			if ok && currentPatLetter != expectedPatLetter {
				return false
			}

			hash[currentPatLetter] = word
			hash2[word] = currentPatLetter
			start = i + 1
			c++
		}
	}

	return c == len(pattern)
}
