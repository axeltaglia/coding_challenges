package main

import "fmt"

func main() {
	fmt.Println(isValid("{["))
}

func isValid(s string) bool {
	st := make([]rune, 0)

	var r bool
	for _, ch := range s {
		st, r = handleSt(st, ch)
		if r {
			return false
		}
	}

	return len(st) == 0
}

func handleSt(st []rune, ch rune) ([]rune, bool) {
	openers := []rune{'(', '[', '{'}
	if inSlice(ch, openers) {
		st = append(st, ch)
		return st, false
	}

	closers := []rune{')', ']', '}'}

	if len(st) > 0 {
		lastElem := st[len(st)-1]
		t := giveMeType(ch, closers)

		if giveMeType(lastElem, openers) != t {
			return nil, true
		}
	} else {
		return nil, true
	}

	st = st[:len(st)-1]

	return st, false
}

func inSlice(e rune, s []rune) bool {
	for _, ch := range s {
		if ch == e {
			return true
		}
	}

	return false
}

func giveMeType(e rune, l []rune) int {
	for i, ch := range l {
		if ch == e {
			return i
		}
	}
	return -1
}
