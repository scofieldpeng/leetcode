// problem: https://leetcode.com/problems/rank-transform-of-an-array/
package arrayRankTransform

import "sort"

func arrayRankTransform(arr []int) []int {
	arrLen := len(arr)
	arrCopy := make([]int, arrLen, arrLen)
	for k, v := range arr {
		arrCopy[k] = v
	}
	sort.Ints(arrCopy)
	arrMap := make(map[int]int)
	start := 1
	for _, v := range arrCopy {
		if _, ok := arrMap[v]; !ok {
			arrMap[v] = start
			start++
		}
	}

	res := make([]int, arrLen, arrLen)
	for k, v := range arr {
		res[k] = arrMap[v]
	}

	return res
}
