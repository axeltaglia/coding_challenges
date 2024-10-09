package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(groupAnagrams([]string{"cab", "tin", "pew", "duh", "may", "ill", "buy", "bar", "max", "doc"}))
}
func groupAnagrams(strs []string) [][]string {
	hash := make(map[string][]string)
	for _, s := range strs {
		sortedString := sortString(s)
		hash[sortedString] = append(hash[sortedString], s)
	}

	output := make([][]string, 0)
	for _, strArr := range hash {
		output = append(output, strArr)
	}
	return output
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
