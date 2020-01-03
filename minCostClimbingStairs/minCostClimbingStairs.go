// problem: https://leetcode.com/problems/min-cost-climbing-stairs
// solution: https://algs.home.pjf.im/min-cost-climbing-stairs.html
package minCostClimbingStairs

// 貌似看起来是动态规划？
// 只要确保每一步都是最小的cost，那么总数也是最小的了，也就是 const += min(cost[i],cost[i+1])
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length, length)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < length; i++ {
		dp[i] = min(cost[i]+dp[i-1], cost[i]+dp[i-2])
	}

	return min(dp[length-1], dp[length-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
