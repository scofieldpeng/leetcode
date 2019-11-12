// problem: https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii
// solution: https://algs.home.pjf.im/best-time-to-buy-and-sell-stock-ii.html
package bestTimeToBuyAndSellStockII

func maxProfit(prices []int) int {
	var (
		currentProfit = 0
		length        = len(prices)
	)

	for i := 1; i < length; i++ {
		// 如果涨价继续持有
		if prices[i]-prices[i-1] > 0 {
			currentProfit += prices[i] - prices[i-1]
		}
	}

	return currentProfit
}
