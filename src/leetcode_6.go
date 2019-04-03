package src

import "fmt"

//算法复杂度为O(n*r)，通过r次循环遍历s，计算出每个字符是否符于当前r行，是则加入到输出字符串里，不是则跳过
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	output := ""

	l := 0
	n := 2*numRows - 2
	for l < numRows {
		for i, _ := range s {
			k := i % n
			if k < numRows {
				if l == k {
					output = output + s[i:i+1]
				}
			} else {
				if l == (n - k) {
					output = output + s[i:i+1]
				}
			}
		}
		l = l + 1
	}
	return output
}

//算法复杂度为O(n)，通过r次循环，通过公式计算出属于r行的字符idx,直接加入到输出字符串里，减少无效计算
func convert_best(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	output := ""

	r := 0
	l := len(s)
	n := 2*numRows - 2
	cols := l / n

	for r < numRows {
		v := 0
		c := 0
		for c <= cols {
			v = c*n + r
			if v < l {
				output = output + s[v:v+1]
				fmt.Println(output)
			}

			if r != 0 && r != numRows-1 {
				v = c*n + n - r
				if v < l {
					output = output + s[v:v+1]
					fmt.Println(output)
				}
			}
			c = c + 1
		}
		r = r + 1
	}
	return output
}
func main() {
	s := "PAYPALISHIRING"
	numRows := 4

	fmt.Println(convert_best(s, numRows) == "PINALSIGYAHRPI")
}
