package isMonotonic

// 这到题目核心就是要找到这个人到底是在增长还是下降
func isMonotonic(A []int) bool {
	length := len(A)
	if length < 3 {
		return true
	}

	isFind := false
	isIncrease := false

	for i := 1; i < length; i++ {
		if A[i] == A[i-1] {
			continue
		}
		// 先确定增减关系
		if !isFind {
			if A[i] > A[i-1] {
				isIncrease = true
			}
			isFind = true
			continue
		}

		// 当出现不一致的时候说明不满足条件
		if (isIncrease && A[i] < A[i-1]) || (!isIncrease && A[i] > A[i-1]) {
			return false
		}
	}

	return true
}
