// problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// solution: https://algs.home.pjf.im/best-time-to-buy-and-sell-stock.html
package bestTimeToBuyAndSellStock

import "math"

func maxProfit(prices []int) int {
	buyPrice := math.MaxInt64
	profit := 0
	for _, v := range prices {
		// 找到最低的买入价格
		if buyPrice > v {
			buyPrice = v
			continue
		}

		if v-buyPrice > profit {
			profit = v - buyPrice
		}
	}

	return profit
}

// dp solution
func maxProfit2(prices []int) int {
	var (
		maxCurrent = 0
		maxSofar   = 0
	)

	for i := 1; i < len(prices); i++ {
		maxCurrent += prices[i] - prices[i-1]
		if maxCurrent <= 0 {
			maxCurrent = 0
		}

		if maxCurrent > maxSofar {
			maxSofar = maxCurrent
		}
	}

	return maxSofar
}