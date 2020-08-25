// problem: https://leetcode.com/problems/minimum-absolute-difference/
package minimumAbsDifference

import (
	"math"
	"sort"
)

// 首先将数组按照大小排序
// 循环遍历，一边找最小差值，一边放置结果
func minimumAbsDifference(arr []int) [][]int {
	if len(arr) < 1 {
		return [][]int{}
	}

	// 先排序
	sort.Ints(arr)

	length := len(arr)
	min := math.MaxInt64
	res := make([][]int, 0)

	// 因为排好序以后，相邻的两个值的绝对差肯定是最小的（比如，[1,2,3]中1,2肯定比1,3小，所以不需要子循环
	for i := 0; i <= length-2; i++ {
		tmp := arr[i+1] - arr[i]
		if tmp < min {
			res = make([][]int, 0)
		}
		if tmp <= min {
			min = tmp
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}

	return res
}
