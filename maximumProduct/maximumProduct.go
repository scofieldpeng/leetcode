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

	return nums[length-1] * nums[length-2] * nums[length-3]
}
