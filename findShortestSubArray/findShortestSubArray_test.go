package findShortestSubArray

import "testing"

func Test_findShortestSubArray(t *testing.T) {
	inputs := [][]int{
		{1, 2, 2, 3, 1},
		{1, 2, 2, 3, 1, 4, 2},
	}
	outputs := []int{
		2, 6,
	}

	for k, v := range inputs {
		res := findShortestSubArray(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
