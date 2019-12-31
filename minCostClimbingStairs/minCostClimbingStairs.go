// problem: https://leetcode.com/problems/min-cost-climbing-stairs
// solution: https://algs.home.pjf.im/min-cost-climbing-stairs.html
package minCostClimbingStairs

// 貌似看起来是动态规划？
// 只要确保每一步都是最小的cost，那么总数也是最小的了，也就是 nextCost = min(cost[i],cost[i+1])
// TODO: 代码看起来还有些细节要注意
func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	totalCost := 0
	for i := -1; i < length-1; {
		if cost[i+1] > cost[i+2] {
			i += 2
		} else {
			i += 1
		}
		totalCost += cost[i]
	}

	return totalCost
}
