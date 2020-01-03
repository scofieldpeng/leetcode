// problem: https://leetcode.com/problems/largest-number-at-least-twice-of-others/
// solution: https://algs.home.pjf.im/largest-number-at-least-twice-of-others.html
package dominantIndex

import "math"

func dominantIndex(nums []int) int {
	// find the max
	maxIndex := -1
	max := math.MinInt64
	for k, v := range nums {
		if v >= max {
			maxIndex = k
			max = v
			continue
		}
	}
	if maxIndex > -1 {
		for k, v := range nums {
			if maxIndex == k {
				continue
			}

			if v*2 > max {
				return -1
			}
		}
	}

	return maxIndex
}
