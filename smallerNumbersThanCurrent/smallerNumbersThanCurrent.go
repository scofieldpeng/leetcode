package smallerNumbersThanCurrent

import "sort"

func smallerNumbersThanCurrent(nums []int) []int {
	length := len(nums)

	numsCopy := make([]int, length, length)
	for k, v := range nums {
		numsCopy[k] = v
	}

	// 排序后能查看出当前的数值
	numNumbers := make(map[int]int)
	sort.Ints(numsCopy)
	for k, v := range numsCopy {
		if _, ok := numNumbers[v]; ok {
			continue
		}
		numNumbers[v] = k
	}

	res := make([]int, length, length)
	for k, v := range nums {
		res[k] = numNumbers[v]
	}

	return res
}
