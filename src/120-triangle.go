package src

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}

	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// compare triangle[i+1][j] triangle[i+1][j+1]
			min := triangle[i+1][j]
			if triangle[i+1][j+1] < min {
				min = triangle[i+1][j+1]
			}
			triangle[i][j] += min
		}
	}
	return triangle[0][0]
}

func ResultOfminimumTotal(triangle [][]int) int {
	return minimumTotal(triangle)
}
