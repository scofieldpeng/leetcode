package kWeakestRows

import "testing"

func Test_kWeakestRows(t *testing.T) {
	inputs1 := [][][]int{
		{
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 0},
			{1, 0, 0, 0, 0},
			{1, 1, 0, 0, 0},
			{1, 1, 1, 1, 1},
		},
		{
			{1, 0, 0, 0},
			{1, 1, 1, 1},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
		{
			{1, 0}, {1, 0}, {1, 0}, {1, 1},
		},
	}
	inputs2 := []int{
		3, 2, 4,
	}
	outputs := [][]int{
		{2, 0, 3},
		{0, 2},
		{0, 1, 2, 3},
	}

	for k, input1 := range inputs1 {
		res := kWeakestRows(input1, inputs2[k])
		if !isEqual(res, outputs[k]) {
			t.Fatalf("\ninput: %v,%v\nwanted: %v\n get: %v\n", input1, inputs2[k], outputs[k], res)
		}
	}
}

func isEqual(a, b []int) bool {
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
