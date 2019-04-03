package src

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	o := 0
	y := x
	for y > 0 {
		o = o*10 + y%10
		y = y / 10
	}

	return x == o
}

func main() {
	fmt.Println(isPalindrome(10))
}
