package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	i := 0
	s := len(strs)

	prefix := ""
	if len(strs) > 0 {
		prefix = strs[0]
	}
	ss := len(prefix)

	for i < ss {
		same := true
		j := 0
		for j < s {
			if len(strs[j]) <= i {
				same = false
				break
			}
			if strs[j][i] != prefix[i] {
				same = false
				break
			}
			j = j + 1
		}
		if !same {
			break
		}
		i = i + 1
	}
	return prefix[0:i]
}

func main() {
	fmt.Printf("longestCommonPrefix is:%s\n", longestCommonPrefix([]string{"", "", ""}))
}
