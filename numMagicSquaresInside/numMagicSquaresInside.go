package numMagicSquaresInside

// 暴力破解
func numMagicSquaresInside(grid [][]int) int {
	if len(grid) < 3 || len(grid[0]) < 3 {
		return 0
	}

	rowLength := len(grid[0])
	columnLength := len(grid)

	find := 0
	for column := 0; column+2 < columnLength; column++ {
		for row := 0; row+2 < rowLength; row++ {
			// 找到起点
			if isMagicSquare(grid, column, row) {
				find++
			}
		}
	}

	return find
}

// 暴力破解查找
func isMagicSquare(grid [][]int, column, row int) bool {
	// 中间不为5肯定不是
	if grid[column+1][row+1] != 5 {
		return false
	}
	num := make([]int, 9, 9)
	for i := column; i < 3+column; i++ {
		for j := row; j < 3+row; j++ {
			// 数字必须为1-9，且只能出现一次
			if grid[i][j] > 9 || grid[i][j] < 1 {
				return false
			}
			num[grid[i][j]-1]++
			if num[grid[i][j]-1] > 1 {
				return false
			}

			if j == row {
				// 右
				if 15 != grid[i][j]+grid[i][j+1]+grid[i][j+2] {
					return false
				}
				// 右上
				if i == 2+column {
					if 15 != grid[i][j]+grid[i-1][j+1]+grid[i-2][j+2] {
						return false
					}
				}
			}

			// 下
			if i == column {
				if 15 != grid[i][j]+grid[i+1][j]+grid[i+2][j] {
					return false
				}

				// 右下
				if j == row {
					if 15 != grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] {
						return false
					}
				}
			}
		}
	}

	return true
}
