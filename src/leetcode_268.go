package src

func missingNumber(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum = sum + n
	}

	s := len(nums)
	rsum := (1 + s) * s / 2
	return rsum - sum
}
