package test

import "testing"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	end := len(nums) - 1
	start := 0

	for {
		for end >= start {
			if nums[end] == val {
				end--
			} else {
				break
			}
		} // nums[end] != value

		for start < end {
			if nums[start] != val {
				start++
			} else {
				break
			}
		} // nums[start] == value
		if end <= start {
			break
		}

		tmp := nums[end]
		nums[start] = tmp
		nums[end] = val
	}

	return end + 1
}

func TestRemoveElement(t *testing.T) {
	a := []int{0, 1, 2, 2, 3, 0, 4, 2}
	t.Log(removeElement(a, 2))
	t.Log(a)
}
