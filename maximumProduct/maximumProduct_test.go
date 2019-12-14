package maximumProduct

import "testing"

func Test_maximumProduct(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3},
		{1, 2},
		{1, 2, 3, 4},
		{5, 2, 1, 2},
		{-6, -5, 1, 2, 5},
	}
	outputs := []int{
		6, 0, 24, 20, 150,
	}

	for k, v := range inputs {
		res := maximumProduct(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
