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
