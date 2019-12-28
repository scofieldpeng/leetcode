// problem: https://leetcode.com/problems/longest-continuous-increasing-subsequence/
// solution: https://algs.home.pjf.im/longest-continuous-increasing-subsequence.html
package findLengthOfLCIS

func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}

	longest := 1
	current := 1

	for i := 1; i < length; i++ {
		// 跳过了
		if nums[i-1] >= nums[i] {
			if longest < current {
				longest = current
			}
			current = 1
			continue
		}
		current++
		if current > longest {
			longest = current
		}
	}

	return longest
}

// 滑动窗口算法
// 这里的核心算法就是，找到开始和结束的下标，然后得出长度
// 每次长度不一样的时候就比对是否是最长的
func findLengthOfLCIS2(nums []int) int {
	longest, start := 0, 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if i > 0 && nums[i-1] >= nums[i] {
			start = i
		}
		longest = max(longest, i-start+1)
	}

	return longest
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
