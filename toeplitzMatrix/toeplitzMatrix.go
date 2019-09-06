// A matrix is Toeplitz if every diagonal from top-left to bottom-right has the same element.
//
//Now given an M x N matrix, return True if and only if the matrix is Toeplitz.
//
//
//Example 1:
//
//Input:
//matrix = [
//  [1,2,3,4],
//  [5,1,2,3],
//  [9,5,1,2]
//]
//Output: True
//Explanation:
//In the above grid, the diagonals are:
//"[9]", "[5, 5]", "[1, 1, 1]", "[2, 2, 2]", "[3, 3]", "[4]".
//In each diagonal all elements are the same, so the answer is True.
//Example 2:
//
//Input:
//matrix = [
//  [1,2],
//  [2,2]
//]
//Output: False
//Explanation:
//The diagonal "[1, 2]" has different elements.
//
//Note:
//
//matrix will be a 2D array of integers.
//matrix will have a number of rows and columns in range [1, 20].
//matrix[i][j] will be integers in range [0, 99].
//
//Follow up:
//
//What if the matrix is stored on disk, and the memory is limited such that you can only load at most one row of the matrix into the memory at once?
//What if the matrix is so large that you can only load up a partial row into the memory at once?
package toeplitzMatrix

// toeplitz matrix矩阵通过维基百科可以得知，其实就是需要a(x,y) = b(x+1,y+1)就行，因此就是一个简单的二重for循环即可
func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) < 1 {
		return true
	}

	xLen,yLen := len(matrix),len(matrix[0])

	// 这里对i,j要大于0进行限制，是因为其实我们可以得知，如果二维数组是x，y轴的表示，x，y轴的贴边这一栏是不需要进行判断的，而放在
	// 代码上看就是i，j不为0，所以不必要从0开始遍历
	for i := 1; i < xLen; i++ {
		for j := 1; j < yLen; j++ {
			if matrix[i-1][j-1] != matrix[i][j] {
				return false
			}
		}
	}

	return true
}
