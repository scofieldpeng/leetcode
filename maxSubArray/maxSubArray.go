// problem: https://leetcode.com/problems/maximum-subarray/
// solution: algs.home.pjf.im/maximum-subarray.html
package maxSubArray

func maxSubArray(nums []int) int {
	var (
		maxSum, currentSum, length = nums[0], nums[0], len(nums)
	)

	for i := 1; i < length; i++ {
		if currentSum > 0 {
			currentSum += nums[i]
		} else {
			currentSum = nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}
