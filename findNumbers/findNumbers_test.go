package findNumbers

import "testing"

func Test_findNumbers(t *testing.T) {
	inputs := [][]int{
		{12, 345, 2, 6, 7896},
		{555,901,482,1771},
	}

	outputs := []int{
		2,
		1,
	}

	for i, input := range inputs {
		res := findNumbers(input)
		if res != outputs[i] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", input, outputs[i], res)
		}
	}
}
