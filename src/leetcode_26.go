package src

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var nextPos int = 0
	var startPos int = 0
	var totalSz int = len(nums)
	var currentValue int = nums[startPos]
	for nextPos < totalSz {
		for nextPos < totalSz && currentValue == nums[nextPos] {
			nextPos++
		}
		nums[startPos] = currentValue
		startPos++
		if nextPos >= totalSz {
			break
		}
		currentValue = nums[nextPos]
	}
	return startPos
}
