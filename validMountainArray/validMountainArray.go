// problem: https://leetcode.com/problems/valid-mountain-array/submissions/
// solution: https://algs.home.pjf.im/valid-mountain-array/submissions.html
package validMountainArray

func validMountainArray(A []int) bool {
	length := len(A)
	if length < 1 {
		return false
	}
	var i int
	// 找到峰顶
	for i = 0; i < length-1; i++ {
		if A[i] > A[i+1] {
			break
		}
		if A[i] == A[i+1] {
			return false
		}
	}
	if i <= 0 || i == length-1 {
		return false
	}

	// 开始下降
	for i := i + 1; i < length; i++ {
		if A[i] >= A[i-1] {
			return false
		}
	}

	return true
}
