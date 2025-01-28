package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("agg", "pee"))
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hash := make(map[rune]rune)
	hash2 := make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		s_ch := rune(s[i])
		t_ch := rune(t[i])
		v, ok := hash[s_ch]
		if ok && v != t_ch {
			return false
		}
		v2, ok2 := hash2[t_ch]
		if ok2 && v2 != s_ch {
			return false
		}
		hash[s_ch] = t_ch
		hash2[t_ch] = s_ch
	}

	return true
}
