package src

func GetRow(rowIndex int) []int {
	array := make([]int, rowIndex+1)
	getRowSub(rowIndex, array)
	return array
}

func getRowSub(rowIndex int, array []int) {
	if rowIndex == 0 {
		array[0] = 1
		return
	}
	getRowSub(rowIndex-1, array)

	tmp := 0
	for i := 0; i < rowIndex; i++ {
		x := tmp + array[i]
		tmp = array[i]
		array[i] = x
	}
	array[rowIndex] = 1
}
