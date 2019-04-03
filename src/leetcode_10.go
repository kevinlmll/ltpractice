package src

import "fmt"

func checkPosition(res [][]bool, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	return res[j][i]
}

func isMatch(s string, p string) bool {
	var res [][]bool

	//初始化结果数组
	i := 0
	j := 0
	sLen := len(s) + 1
	pLen := len(p) + 1

	for j < pLen {
		r := make([]bool, sLen, sLen)
		for i < sLen {
			r[i] = false
			i = i + 1
		}
		res = append(res, r)
		j = j + 1
	}

	//按列进行遍历
	i = 0
	for i < sLen {
		j = 0
		for j < pLen {
			r1 := false
			r2 := false
			r3 := false

			if i == 0 && j == 0 {
				res[j][i] = true
				j = j + 1
				continue
			}

			pIdx := j - 1
			sIdx := i - 1

			if pIdx >= 0 && sIdx >= 0 {
				if p[pIdx] == s[sIdx] || p[pIdx] == '.' || (pIdx >= 1 && p[pIdx] == '*' && (p[pIdx-1] == s[sIdx] || p[pIdx-1] == '.')) {
					r1 = checkPosition(res, i-1, j-1)
				}

				if p[pIdx] == '*' && pIdx >= 1 && (p[pIdx-1] == s[sIdx] || p[pIdx-1] == '.') {
					r3 = checkPosition(res, i-1, j)
				}
			}

			if pIdx >= 0 {
				if p[pIdx] == '*' {
					r2 = checkPosition(res, i, j-2)
				}
			}

			res[j][i] = r1 || r2 || r3
			//fmt.Printf("s:%s, p:%s, result:%v\n", s[0:sIdx], p[0:pIdx], res[pIdx][sIdx])
			j = j + 1
		}

		i = i + 1
	}

	for row, v := range res {
		for col, bv := range v {
			fmt.Printf("row:%d, column:%d, s:%s, p:%s, result:%v\n", row, col, s[0:col], p[0:row], bv)
		}
	}

	return res[pLen-1][sLen-1]
}

func main() {
	fmt.Println(isMatch("aaa", ".*"))
}
