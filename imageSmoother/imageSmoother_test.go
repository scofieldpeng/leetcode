package imageSmoother

import "testing"

func Test_imageSmoother(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},
	}

	outputs := [][][]int{
		{
			{0, 0, 0},
			{0, 0, 0},
			{0, 0, 0},
		},
	}

	for k, v := range inputs {
		res := imageSmoother(v)
		if false == equal(res, outputs[k]) {
			t.Fatalf("input:%v, wanted: %v, get: %v \n", v, outputs[k], res)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if len(v) != len(b[k]) {
			return false
		}
		for kk, vv := range v {
			if vv != b[k][kk] {
				return false
			}
		}
	}

	return true
}
