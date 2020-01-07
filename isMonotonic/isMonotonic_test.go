package isMonotonic

import "testing"

func Test_isMonotonic(t *testing.T) {
	inputs := [][]int{
		{1, 2, 2, 3},
		{6, 5, 4, 4},
		{1, 3, 2},
		{1, 2, 4, 5},
		{1, 1, 1},
	}

	outputs := []bool{
		true, true, false, true, true,
	}

	for k, v := range inputs {
		res := isMonotonic(v)
		if res != outputs[k] {
			t.Fatalf("input:%v,wanted:%v,get:%v\n", v, outputs[k], res)
		}
	}
}
