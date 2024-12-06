package main

func lengthOfLastWord(s string) int {
	acumLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			acumLen++
		} else if acumLen > 0 {
			break
		}
	}

	return acumLen
}
