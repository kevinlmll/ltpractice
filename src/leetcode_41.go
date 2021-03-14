package src

func firstMissingPositive(nums []int) int {
	i := 1
	max := len(nums)

	for i <= max {
		for nums[i-1] != i {
			a := nums[i-1]

			//当前值超出范围，可以不用继续计算
			if a > max || a <= 0 {
				break
			}

			b := nums[a-1]
			//当前值与其期待位置上的值已经一致，不需要继续计算
			if a == b {
				break
			}

			nums[a-1] = a
			nums[i-1] = b
		}
		i = i + 1
	}

	i = 1
	for i <= max {
		if nums[i-1] != i {
			break
		}
		i = i + 1
	}
	return i

}
