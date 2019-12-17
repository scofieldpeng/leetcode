package maximumProduct

import (
	"math"
	"sort"
)

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

// 核心算法比较简单，算是排序算法的改进版本，我们知道结果就是最大的三个数字或者是最大的那个数字和最小的两个数字之间，那么我们只需要遍历一下
// 这个数组，找到这几个数字即可
func maximumProduct2(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	min, secondMin, thirdMax, secondMax, max := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32, math.MinInt32
	for _, v := range nums {
		if v <= min {
			min, secondMin = v, min
		} else if v <= secondMin {
			secondMin = v
		}

		if v >= max {
			thirdMax, secondMax, max = secondMax, max, v
		} else if v >= secondMax {
			thirdMax, secondMax = secondMax, v
		} else if v >= thirdMax {
			thirdMax = v
		}
	}

	lastProduct := thirdMax * secondMax * max
	negativeProduct := min * secondMin * max

	if lastProduct > negativeProduct {
		return lastProduct
	}

	return negativeProduct
}
