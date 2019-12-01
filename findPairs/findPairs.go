// problem: https://leetcode.com/problems/k-diff-pairs-in-an-array/
// solution: https://algs.home.pjf.im/k-diff-pairs-in-an-array/
package findPairs

import "fmt"

// 暴力破解
func findPairs(nums []int, k int) int {
	res := 0
	exitPairMap := make(map[string]bool)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				if _, ok := exitPairMap[fmt.Sprintf("%d:%d", nums[i], nums[j])]; !ok {
					if _, ok := exitPairMap[fmt.Sprintf("%d:%d", nums[j], nums[i])]; !ok {
						res++
						exitPairMap[fmt.Sprintf("%d:%d", nums[i], nums[j])] = true
						exitPairMap[fmt.Sprintf("%d:%d", nums[j], nums[i])] = true
					}
				}
			}
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// hash法
func findPairs2(nums []int, k int) int {
	numsMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := numsMap[v]; !ok {
			numsMap[v] = 0
		}
		numsMap[v]++
	}
	res := 0
	for i := range numsMap {
		if k > 0 {
			if _, ok := numsMap[k+i]; ok {
				res++
			}
			continue
		}
		if k == 0 {
			if v, ok := numsMap[i]; ok && v > 1 {
				res++
			}
		}
	}

	return res
}
