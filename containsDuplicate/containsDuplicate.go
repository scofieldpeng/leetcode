// problem: https://leetcode.com/problems/contains-duplicate
// solution: https:/algs.home.pjf.im/contains-duplicate
package containsDuplicate

import (
	"math"
	"sort"
)

func containsDuplicate(nums []int) bool {
	numsMap := make(map[int]int,len(nums))
	for _, v := range nums {
		if _, exist := numsMap[v]; !exist {
			numsMap[v] = 0
		} else {
			return true
		}
		numsMap[v]++
	}

	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	lastNum := math.MaxInt64
	for _, v := range nums {
		if v == lastNum {
			return true
		}
		lastNum = v
	}

	return false
}
