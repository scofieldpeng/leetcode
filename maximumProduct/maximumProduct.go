package maximumProduct

import "sort"

func maximumProduct(nums []int) int {
	for k, v := range nums {
		if v < 1 {
			nums[k] = -v
		}
	}
	sort.Ints(nums)

	if len(nums) < 3 {
		return 0
	}

	length := len(nums)

	// 要不是最后一个乘以第一二个值，要不就是倒数3个值
	lastProduct := nums[length-1] * nums[length-2] * nums[length-3]
	negativeProduct := nums[length-1] * nums[0] * nums[1]

	if lastProduct > negativeProduct {
		return lastProduct
	}

	return negativeProduct
}
