package src

func Generate(numRows int) [][]int {
	res := make([][]int, numRows)
	generateSub(numRows, res)
	return res
}

func generateSub(numRows int, res [][]int) {
	if numRows <= 1 {
		res[numRows-1] = []int{1}
		return
	}
	generateSub(numRows-1, res)
	row := make([]int, numRows)
	lastRow := numRows - 2
	for i := 0; i < numRows-1; i++ {
		if i == 0 {
			row[i] = res[lastRow][i]
		} else {
			row[i] = res[lastRow][i] + res[lastRow][i-1]
		}
	}
	row[numRows-1] = 1
	res[numRows-1] = row
}
