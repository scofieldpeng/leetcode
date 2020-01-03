package dominantIndex

import "testing"

func Test_dominantIndex(t *testing.T) {
	inputs := [][]int{
		{3, 6, 1, 0,},
		{1, 2, 3, 4,},
	}
	outputs := []int{
		1, -1,
	}

	for k, v := range inputs {
		res := dominantIndex(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, wanted: %v, get:%v\n", v, outputs[k], res)
		}
	}
}
