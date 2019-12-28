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
