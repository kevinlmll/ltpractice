package src

func StrStr(haystack string, needle string) int {
	targetLen := len(needle)
	if targetLen == 0 {
		return 0
	}

	compare := func(idx int) bool {
		for j := 0; j < targetLen; j++ {
			// fmt.Printf("cha:%c, cha:%c\n", needle[j], haystack[idx+j])
			if needle[j] != haystack[idx+j] {
				return false
			}
		}
		return true
	}

	srcLen := len(haystack)
	for i, v := range haystack {
		if byte(v) == needle[0] && (srcLen-i) >= targetLen {
			if compare(i) {
				return i
			}
		}
	}
	return -1
}
