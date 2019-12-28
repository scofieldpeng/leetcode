package checkPossibility

import "testing"

func Test_checkPossibility(t *testing.T) {
	inputs := [][]int{
		{4, 2, 3},
		{4, 2, 1},
		{1, 3, 2},
		{3, 4, 2, 3},
		{-1, 4, 2, 3},
	}
	outputs := []bool{
		true, false, true, false, true,
	}

	for k, v := range inputs {
		res := checkPossibility(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}
