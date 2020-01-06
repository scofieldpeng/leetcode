package maxDistToClosest

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	inputs := [][]int{
		{1, 0, 0, 0, 1, 0, 1},
		{1, 0, 0, 0},
		{0, 1, 0},
		{0, 1, 1, 1, 0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0},
	}
	outputs := []int{
		2, 3, 1, 2, 3,
	}

	for k, v := range inputs {
		res := maxDistToClosest(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}
