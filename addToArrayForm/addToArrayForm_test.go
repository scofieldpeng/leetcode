package addToArrayForm

import "testing"

func Test_addToArrayForm(t *testing.T) {
	inputs1 := [][]int{
		{1, 2, 0, 0},
		{2, 7, 4},
		{2, 1, 5},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3},
	}
	inputs2 := []int{
		34, 181, 806, 1, 516,
	}
	outputs := [][]int{
		{1, 2, 3, 4},
		{4, 5, 5},
		{1, 0, 2, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9},
	}

	for k, v := range inputs1 {
		res := addToArrayForm(v, inputs2[k])
		if false == equal(outputs[k], res) {
			t.Fatalf("input:%v,%v, wanted:%v, get:%v\n", v, inputs2[k], outputs[k], res)
		}
	}
}

func Test_addToArrayForm2(t *testing.T) {
	inputs1 := [][]int{
		{1, 2, 0, 0},
		{2, 7, 4},
		{2, 1, 5},
		{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3},
		{0},
		{0},
	}
	inputs2 := []int{
		34, 181, 806, 1, 516, 23, 1000,
	}
	outputs := [][]int{
		{1, 2, 3, 4},
		{4, 5, 5},
		{1, 0, 2, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9},
		{2, 3},
		{1, 0, 0, 0},
	}

	for k, v := range inputs1 {
		res := addToArrayForm2(v, inputs2[k])
		if false == equal(outputs[k], res) {
			t.Fatalf("input:%v,%v, wanted:%v, get:%v\n", v, inputs2[k], outputs[k], res)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}
