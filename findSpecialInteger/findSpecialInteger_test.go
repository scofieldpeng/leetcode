package findSpecialInteger

import "testing"

func Test_findSpecialInteger(t *testing.T) {
	inputs := [][]int{
		{1, 2, 2, 6, 6, 6, 6, 7, 10},
	}
	outputs := []int{
		6,
	}

	for i, input := range inputs {
		res := findSpecialInteger(input)
		if res != outputs[i] {
			t.Fatalf("input:%v, wanted: %v, get: %v\n", input, outputs[i], res)
		}
	}
}
