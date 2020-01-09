package validMountainArray

import "testing"

func Test_validMountainArray(t *testing.T) {
	inputs := [][]int{
		{2, 1},
		{3, 5, 5},
		{0, 3, 2, 1},
		{},
		{2},
	}

	outputs := []bool{
		false, false, true, false,false,
	}

	for k, v := range inputs {
		res := validMountainArray(v)
		if res != outputs[k] {
			t.Fatalf("input:%v,wanted:%v,get:%v\n", v, outputs[k], res)
		}
	}
}
