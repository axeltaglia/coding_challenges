package main

func canConstruct(ransomNote string, magazine string) bool {
	counts := [26]int{}

	for i := 0; i < len(magazine); i++ {
		counts[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		if counts[ransomNote[i]-'a'] == 0 {
			return false
		}
		counts[ransomNote[i]-'a']--
	}

	return true
}

func canConstructWithHash(ransomNote string, magazine string) bool {
	dic := make(map[rune]int)

	for _, ch := range magazine {
		dic[ch]++
	}

	for _, ch := range ransomNote {
		amount, ok := dic[ch]
		if !ok || amount == 0 {
			return false
		}
		dic[ch]--
	}
	return true
}
