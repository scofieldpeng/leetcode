// problem: https://leetcode.com/problems/k-diff-pairs-in-an-array/
// solution: https://algs.home.pjf.im/k-diff-pairs-in-an-array/
package findPairs

import (
	"fmt"
	"sort"
)

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

// 排序+两点法
// 排序后，大小是一致的，这样就不用暴力运算了，小的绝对在左边，大的都在右边，因为如果i,j的差值小于k，那么说明j太小，如果差值大于k，那么说明i太小了
// 同时需要注意的是对于两个一致的值要避免重复计算
func findPairs3(nums []int, k int) int {
	sort.Ints(nums)
	start := 0
	end := 1

	res := 0

	for ; start < len(nums) && end < len(nums); {
		// 确保同一个值只用了一次
		if (start > 0 && nums[start] == nums[start-1]) || nums[start]+k < nums[end] {
			start++
			// 确保两值不能为一样的
		} else if end <= start || nums[start]+k > nums[end] {
			end++
		} else {
			start++
			res++
		}
	}

	return res
}
