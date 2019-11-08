// problem: https://leetcode.com/problems/pascals-triangle-ii/
// solution: https://algs.home.pjf.im/pascals-triangle-ii/
package generatePascalTriangle2

func getRow(rowIndex int) []int {
	var res = make([]int, rowIndex, rowIndex)
	res[0] = 1
	for i := 1; i < rowIndex; i++ {
		// 第一位永远不会变,所有下标0永远不命中
		for j := i; j >= 1; j-- {
			res[j] = res[j] + res[j-1]
		}
	}

	return res
}
