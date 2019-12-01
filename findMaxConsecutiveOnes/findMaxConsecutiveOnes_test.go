package findMaxConsecutiveOnes

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	inputs := [][]int{
		{1, 1, 0, 1, 1, 1},
		{1, 1, 1, 1, 0, 0, 1},
	}
	outputs := []int{
		3,
		4,
	}

	for k, v := range inputs {
		res := findMaxConsecutiveOnes(v)
		if res != outputs[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}
