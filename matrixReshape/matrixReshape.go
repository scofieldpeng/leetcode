// In MATLAB, there is a very useful function called 'reshape', which can reshape a matrix into a new one with different size but keep its original data.
//
//You're given a matrix represented by a two-dimensional array, and two positive integers r and c representing the row number and column number of the wanted reshaped matrix, respectively.
//
//The reshaped matrix need to be filled with all the elements of the original matrix in the same row-traversing order as they were.
//
//If the 'reshape' operation with given parameters is possible and legal, output the new reshaped matrix; Otherwise, output the original matrix.
//
//Example 1:
//Input:
//nums =
//[[1,2],
// [3,4]]
//r = 1, c = 4
//Output:
//[[1,2,3,4]]
//Explanation:
//The row-traversing of nums is [1,2,3,4]. The new reshaped matrix is a 1 * 4 matrix, fill it row by row by using the previous list.
//Example 2:
//Input:
//nums =
//[[1,2],
// [3,4]]
//r = 2, c = 4
//Output:
//[[1,2],
// [3,4]]
//Explanation:
//There is no way to reshape a 2 * 2 matrix to a 2 * 4 matrix. So output the original matrix.
//Note:
//The height and width of the given matrix is in range [1, 100].
//The given r and c are all positive.
package matrixReshape

// 这道题其实就是一个矩阵的变化，比如从1*4的长方形变为2*2的正方形，或者是变成4*1的长方形，但是没法变成3*2，3*1的形状
// 传入的r就是宽，c就是长，所以判断nums变为一维数组后长度是否是r*c的值即可
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if r < 0 || c < 0 || (len(nums) == r && len(nums[0]) == c) {
		panic("r and c should positive")
	}
	numsOneDimension := make([]int, 0)
	for _, v := range nums {
		numsOneDimension = append(numsOneDimension, v...)
	}
	totalNum := len(nums) * len(nums[0])

	if r*c != totalNum {
		return nums
	}

	res := make([][]int, 0, r)
	for i := 0; i < r; i++ {
		res = append(res, numsOneDimension[i*c:i*c+c])
	}

	return res
}
