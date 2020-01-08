package hasGroupsSizeX

import "testing"

func Test_hasGroupSizeX(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 4, 3, 2, 1},
		{1, 1, 1, 2, 2, 2, 3, 3},
		{1},
		{1, 1},
		{1, 1, 2, 2, 2, 2},
		{1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
	}

	outputs := []bool{
		true, false, false, true, true, true,
	}

	for k, v := range inputs {
		res := hasGroupsSizeX(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}
