package numPairsDivisibleBy60

import "testing"

func Test_numPairsDivisibleBy60(t *testing.T) {
	inputs := [][]int{
		{30, 20, 150, 100, 40},
		{60, 60, 60},
	}
	outputs := []int{
		3, 3,
	}

	for k, v := range inputs {
		res := numPairsDivisibleBy60(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, get:%v ,wanted: %v\n", v, res, outputs[k])
		}
	}
}
