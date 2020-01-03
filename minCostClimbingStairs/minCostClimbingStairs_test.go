package minCostClimbingStairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	inputs := [][]int{
		{10, 15, 20},
		{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
	}
	outputs := []int{
		15, 6,
	}

	for k, v := range inputs {
		res := minCostClimbingStairs(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, get:%v, wanted:%v\n", v, res, outputs[k])
		}
	}
}
