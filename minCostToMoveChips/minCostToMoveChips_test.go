package minCostToMoveChips

import (
	"log"
	"testing"
)

func Test_minCostToMoveChips(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{2, 2, 2, 3, 3},
	}
	outputs := []int{
		1, 2,
	}

	for k, v := range inputs {
		res := minCostToMoveChips(v)
		if res != outputs[k] {
			log.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}
