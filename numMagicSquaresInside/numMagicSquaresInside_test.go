package numMagicSquaresInside

import "testing"

func Test_numMagicSquaresInside(t *testing.T) {
	inputs := [][][]int{
		{
			{4, 3, 8, 4},
			{9, 5, 1, 9},
			{2, 7, 6, 2,},
		},
		{
			{5, 5, 5},
			{5, 5, 5},
			{5, 5, 5},
		},
		{
			{3, 2, 1, 6},
			{5, 9, 6, 8},
			{1, 5, 1, 2},
			{3, 7, 3, 4},
		},
		{
			{7, 6, 2, 2, 4},
			{4, 4, 9, 2, 10},
			{9, 7, 8, 3, 10},
			{8, 1, 9, 7, 5},
			{7, 10, 4, 11, 6},
		},
	}

	outputs := []int{
		1,
		0,
		0,
		0,
		0,
	}

	for k, v := range inputs {
		res := numMagicSquaresInside(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, get:%v ,wanted:%v\n", v, res, outputs[k])
		}
	}
}
