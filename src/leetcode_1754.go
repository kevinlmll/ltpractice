package main

import (
	"fmt"
)

func largestMerge(word1 string, word2 string) string {
	var (
		i, j, l1, l2 int
		result       string
	)

	i = 0
	j = 0
	l1 = len(word1)
	l2 = len(word2)

	for i < l1 && j < l2 {
		r := compareSubString(word1[i], word2[j])
		if r != 0 {
			// not equal
			if r < 0 {
				// l < r
				result = result + string(word2[j])
				//fmt.Printf("append:%v at w2[%d]\n", word2[j], j)
				j++
			} else {
				// l > r
				result = result + string(word1[i])
				//fmt.Printf("append:%v at w1[%d]\n", word1[i], i)
				i++
			}
		} else {
			// equal
			winner := dealStringEqual(word1, word2, i, j)
			if winner == 1 {
				result = result + string(word1[i])
				i = i + 1
			} else if winner == 2 {
				result = result + string(word2[j])
				j = j + 1
			}
		}
		//fmt.Printf("result:%v, i:%d, j:%d\n", result, i, j)
	}
	if i < l1 {
		result = result + word1[i:]
	}
	if j < l2 {
		result = result + word2[j:]
	}
	//fmt.Printf("result:%v, i:%d, j:%d\n", result, i, j)

	return result
}

func dealStringEqual(m, n string, i, j int) (winner int) {
	var (
		ii, jj int
		lm, ln int
	)
	lm = len(m)
	ln = len(n)
	ii = i
	jj = j

	for ii < lm && jj < ln {
		if m[ii] != n[jj] {
			break
		}
		ii++
		jj++
	}

	if ii == lm {
		winner = 2
		return
	}
	if jj == ln {
		winner = 1
		return
	}

	if m[ii] > n[jj] {
		winner = 1
	} else {
		winner = 2
	}
	return
}

func compareSubString(l1, l2 uint8) int {
	if l1 == l2 {
		return 0
	}

	if l1 < l2 {
		return -1
	}
	return 1
}

func main() {
	a := []string{
		"jvjjjjjjvjjvjvjjjvjvjjjj",
	}
	b := []string{
		"jjjjjjvjjjjjjvjjjjv",
	}
	result := []string{
		"jvjjjjjjvjjvjvjjjvjvjjjjjjvjjjjjjvjjjjvjjjj",
	}

	for i := 0; i < len(a); i++ {
		r := largestMerge(a[i], b[i])
		fmt.Printf("equal:%v\n", result[i] == r)
	}
}
