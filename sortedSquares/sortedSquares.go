package sortedSquares

import (
	"sort"
)

func sortedSquares(A []int) []int {
	var res = make([]int, 0, len(A))
	for _, v := range A {
		res = append(res, v*v)
	}

	sort.Ints(res)

	return res
}
