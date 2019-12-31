package pivotIndex

import "testing"

func Test_pivotIndex(t *testing.T) {
	inputs := [][]int{
		{1, 7, 3, 6, 5, 6},
		{1, 2, 3},
	}
	outputs := []int{
		3, -1,
	}

	for k, v := range inputs {
		res := pivotIndex(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}
