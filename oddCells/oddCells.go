// problem: https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/
package oddCells

// 这道题目最开始没看懂，鄙视下自己的英语能力
// 其实只需要循环一遍，然后按照规律给矩阵增值就好了，然后第二次循环看下奇数多少个
func oddCells(n int, m int, indices [][]int) int {
	res := 0
	matrix := make([][]int, n, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m, m)
	}

	for _, v := range indices {
		ri, ci := v[0], v[1]
		// 在row上增加值
		for rowK := range matrix[ri] {
			matrix[ri][rowK]++
		}
		// 在column增加值
		for rowIndex := 0; rowIndex < n; rowIndex++ {
			matrix[rowIndex][ci]++
		}
	}

	for _, v := range matrix {
		for _, vv := range v {
			if vv%2 != 0 {
				res++
			}
		}
	}

	return res
}

// 第二种可以优化一下，第二次遍历可以去掉，第一次增值的时候其实就应能找到结果了
// 当给row增加的时候，可以看一下，如果该列不是接下来要增加的列，那么就可以直接判断奇偶数了
// todo
func oddCells2(n int, m int, indices [][]int) int {
	res := 0
	rowMap := make(map[int]int)
	colMap := make(map[int]int)

	for _, v := range indices {
		ri, ci := v[0], v[1]
		if _, ok := rowMap[ri]; !ok {
			rowMap[ri] = 0
		}
		rowMap[ri]++
		// 如果是奇数，说明值也是奇数，一口气增加m数值
		if rowMap[ri]%2 != 0 {
			res += m
		} else {
			res -= m
		}

		// 竖线的就有点意思了，看一下会碰到几根横线，同时和横线的数值相比较
		//
		if _, ok := colMap[ci]; !ok {
			colMap[ci] = 0
		}
		colMap[ci]++
		if colMap[ci]%2 != 0 {
			res += n
		} else {
			res -= n
		}
	}

	return res
}
