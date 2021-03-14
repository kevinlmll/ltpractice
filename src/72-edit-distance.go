package src

func minDistance(word1 string, word2 string) int {
	m := len(word1) + 1
	n := len(word2) + 1
	if m == 0 || n == 0 {
		return m | n
	}

	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		matrix[i][0] = i
	}
	for i := 0; i < n; i++ {
		matrix[0][i] = i
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				a := matrix[i-1][j-1]
				b := matrix[i-1][j]
				c := matrix[i][j-1]
				min := a
				if b < min {
					min = b
				}
				if c < min {
					min = c
				}
				matrix[i][j] = min + 1
			}
		}
	}
	return matrix[m-1][n-1]
}

func ResultOfminDistance(word1, word2 string) int {
	return minDistance(word1, word2)
}
