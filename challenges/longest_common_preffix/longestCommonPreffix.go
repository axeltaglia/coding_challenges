package main

func longestCommonPrefix(strs []string) string {
	var commonPrefix string

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	for {
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 || i > len(strs[j-1])-1 || strs[j][i] != strs[j-1][i] {
				return commonPrefix
			}
		}
		commonPrefix += string(strs[0][i])
		i++
	}
}
