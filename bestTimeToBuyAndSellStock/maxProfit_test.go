package bestTimeToBuyAndSellStock

import "testing"

func Test_maxProfit(t *testing.T) {
	inputs := [][]int{
		{7, 1, 5, 3, 6, 4},
		{7, 6, 4, 3, 1},
		{2, 4, 1},
	}
	outputs := []int{
		5, 0, 2,
	}

	for k, v := range inputs {
		res := maxProfit(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
