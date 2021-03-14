package src

func getDistance(a, b int) int {
	var neg int = 0

	if a < 0 {
		neg = neg + 1
	}

	if b < 0 {
		neg = neg + 1
	}

	if neg == 1 {
		if a < 0 {
			return b - a
		}
		return a - b
	}

	l := a - b
	if l < 0 {
		return -l
	}
	return l
}

//O(n^3)
func threeSumClosest(nums []int, target int) int {
	var size int = len(nums)
	var min int = 100000000
	var score int = 0

	distances := make([][]int, size)
	for i, _ := range distances {
		distances[i] = make([]int, size)
	}

	var i int = 0
	var j int = 0

	for i < size {
		j = 0
		for j < size {
			if i == j {
				distances[i][j] = nums[i]
			} else {
				distances[i][j] = nums[i] + nums[j]
			}
			//fmt.Printf("pos at(%d,%d), sum:%d\n", i, j, distances[i][j])
			j = j + 1
		}
		i = i + 1
	}

	var k int = 0
	i = 0
	for i < size {
		j = 0
		for j < size {
			k = 0
			for k < size {
				if i != j && i != k && j != k {
					allsum := distances[i][j] + nums[k]
					d := getDistance(allsum, target)
					if d < min {
						min = d
						score = allsum
					}
					//fmt.Printf("pos at(%d,%d,%d), sum:%d, distance:%d\n", i, j, k, allsum, d)
				}

				k = k + 1
			}
			j = j + 1
		}
		i = i + 1
	}

	return score
}
