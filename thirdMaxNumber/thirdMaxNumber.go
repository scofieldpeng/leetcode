package thirdMaxNumber

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)

	if len(nums) > 2 {
		return nums[len(nums)-3]
	}

	return nums[len(nums)-1]
}
