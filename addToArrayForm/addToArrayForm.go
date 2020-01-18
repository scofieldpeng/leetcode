package addToArrayForm

// 为了防止整数范围溢出，不能将数组转化为int然后算完后再转化回来，因此只能将整数转化为array
// 我们将A的最高位开始和K相加，如果相加后大于10，那么个位数为当前的值，处于10的值就是下一次参与计算的值
func addToArrayForm(A []int, K int) []int {
	res := make([]int, 0)
	length := len(A)

	for i := length - 1; K > 0 || i >= 0; i-- {
		kv := 0
		// 得到k当前位数
		if K >= 10 {
			kv = K % 10
			K = K / 10
		} else {
			kv = K
			K = 0
		}
		if i >= 0 {
			kv = kv + A[i]
		}
		// 如果k的当前位数和a当前位数大于2位数，那么下一位+1
		if kv >= 10 {
			K += 1
			kv = kv % 10
		}
		res = append(res, kv)
	}

	// 翻转数组
	length = len(res)
	for i := 0; i < length/2; i++ {
		res[i], res[length-i-1] = res[length-i-1], res[i]

	}

	return res
}
