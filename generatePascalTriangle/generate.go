// problem: https://leetcode.com/problems/pascals-triangle/
// solution: https://algs.home.pjf.im/pascals-triangle.html
package generatePascalTriangle

func generate(numRows int) [][]int {
	var (
		res = make([][]int, 0, numRows)
	)

	for n := 1; n <= numRows; n++ {
		tmp := make([]int, 0, n)
		for m := 0; m < n; m++ {
			v := 1
			if m > 0 && m < n-1 {
				v = res[n-2][m-1] + res[n-2][m]
			}
			tmp = append(tmp, v)
		}
		res = append(res, tmp)
	}

	return res
}
