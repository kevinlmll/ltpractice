package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, exclude1, exclude2, target int) [][]int {
	//nums已经排好序
	if len(nums)-2 < 2 {
		return nil
	}

	i := 0
	j := len(nums) - 1
	output := make([][]int, 0)

	for i < j {
		for i == exclude1 || i == exclude2 {
			i = i + 1
		}
		for j == exclude1 || j == exclude2 {
			j = j - 1
		}

		if i >= j {
			break
		}

		sum := nums[i] + nums[j]
		if sum > target {
			j = j - 1
		} else if sum < target {
			i = i + 1
		} else {
			r := []int{nums[i], nums[j], nums[exclude1], nums[exclude2]}
			sort.Ints(r)
			output = append(output, r)
			i = i + 1
		}
	}
	return output
}

func fourSum(nums []int, target int) [][]int {
	output := make([][]int, 0)
	fourSumArray := make([][]int, 0)
	i := 0
	j := 0

	if len(nums) < 4 {
		return output
	}

	sort.Ints(nums)
	for i < len(nums) {
		j = i + 1
		for j < len(nums) {
			r := twoSum(nums, i, j, target-nums[i]-nums[j])
			if r != nil && len(r) > 0 {
				fourSumArray = append(fourSumArray, r...)
			}
			j = j + 1
		}
		i = i + 1
	}

	existArray := make(map[string]int, 0)
	for _, a := range fourSumArray {
		key := fmt.Sprintf("%d_%d_%d_%d", a[0], a[1], a[2], a[3])
		if _, e := existArray[key]; !e {
			existArray[key] = 1
			output = append(output, a)
		}
	}

	return output
}

func main() {
	fmt.Printf("fourSum return:%+v\n", fourSum([]int{-3, -1, 0, 2, 4, 5}, 0))
}
