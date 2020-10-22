// problem: https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/
package countNegatives

// 找到最开始负数的行和列
func countNegatives(grid [][]int) int {
	res := 0
	isColumnNegative := false
	for _, v := range grid {
		// 整行都是负数
		if isColumnNegative {
			res += len(v)
			continue
		}
		for kk, vv := range v {
			// 找到一列中负数的数量
			if vv < 0 {
				// 第一列都是空白，说明下面的所有行都是空白
				if kk == 0 {
					isColumnNegative = true
				}
				res += len(v) - kk
				break
			}
		}
	}

	return res
}
