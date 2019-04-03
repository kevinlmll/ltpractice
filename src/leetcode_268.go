package src

import "fmt"

func missingNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}

	s := len(nums)
	rsum := (1 + s) * s / 2
	return rsum - sum
}

func main() {
	fmt.Printf("missingNumber is:%d\n", missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
