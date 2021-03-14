package src

func minPathSum(grid [][]int) int {
	// rows, columns >= 1
	rows := len(grid)
	columns := len(grid[0])

	minSums := make([]int, columns)
	for r := rows - 1; r >= 0; r-- {
		for c := columns - 1; c >= 0; c-- {
			if r == (rows-1) && c == (columns-1) {
				minSums[c] = grid[r][c]
			} else {
				rightMin := valueAtPosition(minSums, r, c+1, rows, columns)
				bottomMin := valueAtPosition(minSums, r+1, c, rows, columns)
				if rightMin < bottomMin {
					minSums[c] = rightMin
				} else {
					minSums[c] = bottomMin
				}
				minSums[c] += grid[r][c]
			}
		}
	}
	return minSums[0]
}

func valueAtPosition(sums []int, r, c, rows, columns int) int {
	if r >= rows || c >= columns {
		return 20201
	}
	if r < 0 || c < 0 {
		return 20201
	}
	return sums[c]
}

func ResultOfminPathSum(grid [][]int) int {
	return minPathSum(grid)
}
