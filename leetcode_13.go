package main

import "fmt"

func romanToInt(s string) int {
	num := 0
	l := len(s)
	i := 0

	for i = 0; i < l; {
		switch s[i] {
		case 'I':
			//有可能是1，4，9
			if i+1 < l {
				if s[i+1] == 'V' {
					num = num + 4
					i = i + 1
				} else if s[i+1] == 'X' {
					num = num + 9
					i = i + 1
				} else {
					num = num + 1
				}
			} else {
				num = num + 1
			}
		case 'V':
			num = num + 5
		case 'X':
			//有可能是10，40，90
			if i+1 < l {
				if s[i+1] == 'L' {
					num = num + 40
					i = i + 1
				} else if s[i+1] == 'C' {
					num = num + 90
					i = i + 1
				} else {
					num = num + 10
				}
			} else {
				num = num + 10
			}
		case 'L':
			num = num + 50
		case 'C':
			//有可能是100，400，900
			if i+1 < l {
				if s[i+1] == 'D' {
					num = num + 400
					i = i + 1
				} else if s[i+1] == 'M' {
					num = num + 900
					i = i + 1
				} else {
					num = num + 100
				}
			} else {
				num = num + 100
			}
		case 'D':
			num = num + 500
		case 'M':
			num = num + 1000
		}
		i = i + 1
	}
	return num
}

func main() {
	fmt.Printf("romanToInt:%d", romanToInt("CD"))
}
