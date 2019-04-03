package main

import "fmt"

func reverse(x int) int {
	if x == 0 {
		return x
	}

	sign := 1
	rx := x
	if x < 0 {
		sign = -1
		rx = sign * rx
	}

	//rx是正数
	mod := rx % 10
	rx = rx / 10
	for mod == 0 && rx > 0 {
		mod = rx % 10
		rx = rx / 10
	}

	//此时mod不为0
	rx = rx*10 + mod
	o := 0
	for rx > 0 {
		o = o*10 + (rx % 10)
		rx = rx / 10
	}

	o = o * sign
	//max := 2147483647
	max := (1 << 31) - 1
	//min := -2147483648
	min := -(1 << 31)
	fmt.Println(min, max)
	if o >= min && o <= max {
		return o
	}

	return 0
}

func main() {
	fmt.Println(reverse(123))
}
