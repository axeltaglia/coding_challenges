package main

import "fmt"

func main() {
	//fmt.Println(isAnagram("rioir", "oiirr"))
	// func2()
	detectType()
}

func func2() {
	for _, ch := range "Hóla" {
		fmt.Printf("%#U\n", ch)
	}

	fmt.Println("")

	s := "Hóla"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}

	fmt.Println("")

	sliceByte := []byte("Hóla")
	for _, b := range sliceByte {
		fmt.Printf("%c\n", b)
	}

	fmt.Println("")

}

func returnAny() interface{} {
	return 3.23
}

func detectType() {
	r := returnAny()

	switch r.(type) {
	case int:
		fmt.Printf("It's an integer %d\n", r)
	case string:
		fmt.Printf("It's an string %s\n", r)
	default:
		fmt.Printf("undefined type %v\n", r)
	}
}

func isAnagram(s string, t string) bool {
	hash := make(map[rune]int)

	for _, ch := range s {
		hash[ch]++
	}

	for _, ch := range t {
		hash[ch]--
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}
