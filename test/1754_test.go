package test

import (
	"fmt"
	"strings"
	"testing"
)

func largestMerge(word1 string, word2 string) string {
	var (
		i, j, l1, l2 int
		result       string
	)

	w1secs := incrSubStrings(word1)
	//fmt.Printf("incr substrs-1:%v\n", w1secs)
	w2secs := incrSubStrings(word2)
	//fmt.Printf("incr substrs-2:%v\n", w2secs)

	i = 0
	j = 0
	l1 = len(w1secs)
	l2 = len(w2secs)

	for i < l1 && j < l2 {
		r := compareSubString(w1secs[i], w2secs[j])
		if r != 0 {
			// not equal
			if r < 0 {
				// l < r
				result = result + w2secs[j]
				fmt.Printf("append:%v at w2[%d]\n", w2secs[j], j)
				j++
			} else {
				// l > r
				result = result + w1secs[i]
				fmt.Printf("append:%v at w1[%d]\n", w1secs[i], i)
				i++
			}
		} else {
			// equal
			offset, winner := dealStringEqual(w1secs, w2secs, i, j)
			if winner == 0 {
				result = result + strings.Join(w1secs[i:i+offset], "")
				result = result + strings.Join(w2secs[j:j+offset], "")
				fmt.Printf("append:%v at w1[%d], w2[%d]\n", w1secs[i:i+offset], i, j)
				i = i + offset
				j = j + offset
			} else if winner == 1 {
				result = result + strings.Join(w1secs[i:i+offset], "")
				fmt.Printf("append:%v at w1[%d]\n", w1secs[i:i+offset], i)
				i = i + offset
			} else if winner == 2 {
				result = result + strings.Join(w2secs[j:j+offset], "")
				fmt.Printf("append:%v at w2[%d]\n", w2secs[j:j+offset], j)
				j = j + offset
			}
		}
		//fmt.Printf("result:%v, i:%d, j:%d\n", result, i, j)
	}
	if i >= l1 {
		result = result + strings.Join(w2secs[j:], "")
	}
	if j >= l2 {
		result = result + strings.Join(w1secs[i:], "")
	}

	return result
}

func dealStringEqual(m, n []string, i, j int) (offset, winner int) {
	// 如果两者相等，则找A，B各自的一个上增序列
	istart := i
	iend := i + 1
	for ; iend < len(m); iend++ {
		r := compareSubString(m[iend], m[iend-1])
		if r == -1 {
			break
		}
	} // m[istart:iend]

	jstart := j
	jend := j + 1
	for ; jend < len(n); jend++ {
		r := compareSubString(m[jend], m[jend-1])
		if r == -1 {
			break
		}
	} // n[jstart:jend]

	mincr := m[istart:iend]
	nincr := n[jstart:jend]
	fmt.Printf("m-incr:%v, n-incr:%v\n", mincr, nincr)

	l := iend - istart
	if l > (jend - jstart) {
		l = jend - jstart
	}

	winner = -1
	z := 0
	for z = 0; z < l; z++ {
		r := compareSubString(mincr[istart+z], nincr[jstart+z])
		if r != 0 {
			if r == -1 {
				// m < n
				winner = 2
			} else {
				// m > n
				winner = 1
			}
			break
		}
	}
	offset = z
	if winner != -1 {
		return
	}

	mall := strings.Join(m[i:], "")
	nall := strings.Join(n[j:], "")
	r := compareSubString(mall, nall)
	if r == -1 {
		// m < n
		winner = 2
	} else if r == 1 {
		// m > n
		winner = 1
	}
	winner = 0
	return
}

func incrSubStrings(m string) []string {
	l := len(m)
	strs := make([]string, 0, 3000)
	var i, j int
	i = 0
	j = 1
	max := m[i]
	for j < l {
		if m[j] < max {
			strs = append(strs, m[i:j])
			i = j
			max = m[i]
		} else {
			max = m[j]
		}
		j++
	}
	strs = append(strs, m[i:j])
	return strs
}

func compareSubString(l1, l2 string) int {
	if l1 == l2 {
		return 0
	}

	s1 := l1 + l2
	s2 := l2 + l1
	if s1 < s2 {
		return -1
	} else if s1 > s2 {
		return 1
	}
	return 0
}

func TestLargestMerge(t *testing.T) {
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
		t.Error(result[i] != r)
	}
}
