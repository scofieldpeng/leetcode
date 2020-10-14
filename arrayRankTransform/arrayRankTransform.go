// problem: https://leetcode.com/problems/rank-transform-of-an-array/
package arrayRankTransform

import "sort"

func arrayRankTransform(arr []int) []int {
	maxScore := len(arr)
	res := make([]int, maxScore, maxScore)
	arrMap := make(map[int][]int)
	for k, v := range arr {
		if _, ok := arrMap[v]; !ok {
			arrMap[v] = make([]int, 0)
		}
		arrMap[v] = append(arrMap[v], k)
	}

	sort.Ints(arr)
	rank := 0
	for _, v := range arr {
		// 从1开始
		for kk, vv := range arrMap[v] {
			if res[vv] == 0 {
				if kk == 0 {
					rank++
				}
				res[vv] = rank
			}
		}
	}

	return res
}
