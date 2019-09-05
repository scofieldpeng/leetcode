// Given a matrix A, return the transpose of A.
//
//The transpose of a matrix is the matrix flipped over it's main diagonal, switching the row and column indices of the matrix.
//
//
//
//Example 1:
//
//Input: [[1,2,3],[4,5,6],[7,8,9]]
//Output: [[1,4,7],[2,5,8],[3,6,9]]
//Example 2:
//
//Input: [[1,2,3],[4,5,6]]
//Output: [[1,4],[2,5],[3,6]]
//
//
//Note:
//
//1 <= A.length <= 1000
//1 <= A[0].length <= 1000
package transpose

/// 这个只是简单的考察一下矩阵变化，不难发现，其实就是x,y轴的对掉，也就是data[i][j] = data[j][i]
func transpose(A [][]int) [][]int {
	if len(A) < 1 {
		return nil
	}

	var lenRes = len(A[0])
	var res = make([][]int, lenRes, lenRes)

	for _, row := range A {
		for j, v := range row {
			res[j] = append(res[j], v)
		}
	}

	return res
}
