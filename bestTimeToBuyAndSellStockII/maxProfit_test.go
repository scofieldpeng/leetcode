package bestTimeToBuyAndSellStockII

import "testing"

func Test_maxProfit(t *testing.T) {
	input := [][]int{
		{7, 1, 5, 3, 6, 4},
		{1, 2, 3, 4, 5},
		{5, 4, 3, 2, 1},
	}
	output := []int{
		7, 4, 0,
	}

	for k, v := range input {
		res := maxProfit(v)
		if res != output[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, output[k], res)
		}
	}

}
