package thirdMaxNumber

import (
	"math"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	data := make([]int, 3, 3)
	data[0] = math.MinInt64
	data[1] = math.MinInt64
	data[2] = math.MinInt64
	for _, v := range nums {
		calculateThird(data, v)
	}

	if data[2] == math.MinInt64 {
		return data[0]
	}

	return data[2]
}

func calculateThird(nums []int, data int) {
	if data == nums[0] || data == nums[1] || data == nums[2] {
		return
	}
	if data > nums[2] {
		nums[2], data = data, nums[2]
		if nums[2] > nums[1] {
			nums[1], nums[2] = nums[2], nums[1]
			if nums[1] > nums[0] {
				nums[0], nums[1] = nums[1], nums[0]
			}
		}
	}
}
