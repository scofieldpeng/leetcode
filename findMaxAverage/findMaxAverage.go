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

	// 注意这里的k-1,让其默认值为第一个最大，避免k<len(nums)时的case
	maxValue := numSumArr[k-1]
	for i := k; i < length; i++ {
		v := numSumArr[i] - numSumArr[i-k]
		if maxValue < v {
			maxValue = v
		}
	}

	return float64(maxValue) / float64(k)
}

func findMaxAverage2(nums []int, k int) float64 {
	sumValue := 0
	for i := 0; i < k; i++ {
		sumValue += nums[i]
	}

	// 这里sumValue默认赋值给maxValue是防止k<len(nums)时的边界条件
	maxValue := sumValue
	for i := k; i < len(nums); i++ {
		sumValue = sumValue + nums[i] - nums[i-k]
		if sumValue > maxValue {
			maxValue = sumValue
		}
	}

	return float64(maxValue) / float64(k)

}