package findMaxAverage

import "testing"

func Test_findMaxAverage(t *testing.T) {
	inputs1 := [][]int{
		{1, 12, -5, -6, 50, 3},
		{5},
	}
	inputs2 := []int{
		4,
		1,
	}
	outputs := []float64{
		12.75,
		5,
	}

	for k, v := range inputs1 {
		res := findMaxAverage(v, inputs2[k])
		if res != outputs[k] {
			t.Fatalf("input: %v, %v, wanted: %v, get: %v\n", v, inputs2[k], outputs[k], res)
		}
	}
}

func Test_findMaxAverage2(t *testing.T) {
	inputs1 := [][]int{
		{1, 12, -5, -6, 50, 3},
		{5},
	}
	inputs2 := []int{
		4, 1,
	}
	outputs := []float64{
		12.75,
		5,
	}

	for k, v := range inputs1 {
		res := findMaxAverage2(v, inputs2[k])
		if res != outputs[k] {
			t.Fatalf("input: %v, %v, wanted: %v, get: %v\n", v, inputs2[k], outputs[k], res)
		}
	}
}
