package addToArrayForm

// 为了防止整数范围溢出，不能将数组转化为int然后算完后再转化回来，因此只能将整数转化为array
// 又因为这个东西
func addToArrayForm(A []int, K int) []int {
	// 现将k转化为array
	res := make([]int, 1, 1)

	kv := 0
	for i := len(A) - 1; K > 0 || i >= 0; i-- {
		// K当前的值
		if K > 10 {
			kv = K % 10
			K = K / 10
		} else {
			kv = K
			K = 0
		}

		currentRes := kv + A[i] + res[i]
		if currentRes-10 > 0 {
			currentRes -= 10
			res = append(res, currentRes, 1)
		} else {
			res = append(res, currentRes, 0)
		}
	}

	return res
}
