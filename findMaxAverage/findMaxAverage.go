package findMaxAverage

func findMaxAverage(nums []int, k int) float64 {
	length := len(nums)
	var numSumArr = make([]int, length, length)
	// 遍历整个数组
	// 创造一个新的数组用来存储从0到该序列的sum值
	for i := 0; i < length; i++ {
		if i == 0 {
			numSumArr[i] = nums[i]
			continue
		}
		numSumArr[i] = nums[i] + numSumArr[i-1]
	}

	// 找到k个数组中最大的那个值，那么平均值也是最大的
	// 怎么知道a到b的总和呢？
	// 用b到0的总和减去a到0的总和即可
	maxValue := 0
	for i := k; i < length; i++ {
		v := numSumArr[i] - numSumArr[i-k]
		if maxValue < v {
			maxValue = v
		}
	}

	return float64(maxValue) / float64(k)
}
