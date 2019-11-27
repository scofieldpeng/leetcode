package thirdMaxNumber

import "testing"

func Test_thirdMaxNumber(t *testing.T) {
	inputs := [][]int{
		{3, 2, 1},
		{1, 2},
		{2, 2, 3, 1},
	}
	outputs := []int{
		1, 2, 1,
	}

	for k, v := range inputs {
		res := thirdMax(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
