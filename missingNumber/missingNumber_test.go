package missingNumber

import "testing"

func Test_missingNumber(t *testing.T) {
	inputs := [][]int{
		{3, 0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
		{1},
		{0},
	}
	outputs := []int{
		2, 8, 0, 1,
	}

	for k, v := range inputs {
		res := missingNumber(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}

	for k, v := range inputs {
		res := missingNumber2(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}

	for k, v := range inputs {
		res := missingNumber3(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
