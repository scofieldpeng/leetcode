package majorityElement

import "testing"

func Test_majorityElement(t *testing.T) {
	inputs := [][]int{
		{3, 2, 3},
		{2, 2, 1, 1, 1, 2, 2},
	}
	outputs := []int{
		3, 2,
	}

	for k, v := range inputs {
		res := majorityElement(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}

func Test_majorElementByDivideAndConquer(t *testing.T) {
	inputs := [][]int{
		{3, 2, 3},
	}
	outputs := []int{
		3,
	}

	for k, v := range inputs {
		res := majorElementByDivideAndConquer(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
