// problem: https://leetcode.com/problems/shift-2d-grid/
package shiftGrid

func shiftGrid(grid [][]int, k int) [][]int {
	// 初始化结果集
	rowNum := len(grid)
	columnNum := len(grid[0])
	res := make([][]int, rowNum, rowNum)
	for i := 0; i < rowNum; i++ {
		res[i] = make([]int, columnNum, columnNum)
	}

	// 确定新的要挪动的步数, 其实就是想知道每一位要移动的行和列
	moveNum := k % (len(grid[0]) * len(grid))

	for i := 0; i < rowNum; i++ {
		for j := 0; j < columnNum; j++ {
			// 根据现在grid[i][j]的坐标，算出新的结果的行和列的坐标
			columnIndex := (j + moveNum) % columnNum
			// 这里最后余一下是将最后一位偏移到头部
			rowIndex := ((i*columnNum + j + moveNum) / columnNum ) % rowNum

			res[rowIndex][columnIndex] = grid[i][j]
		}
	}

	return res
}
