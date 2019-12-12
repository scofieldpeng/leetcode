package findUnsortedSubarray

import "testing"

func Test_findUnsortedSubArray(t *testing.T) {
	inputs := [][]int{
		{2, 6, 4, 8, 10, 9, 15},
	}
	outputs := []int{
		5,
	}

	for k, v := range inputs {
		res := findUnsortedSubarray(v)
		if res != outputs[k] {
			t.Fatalf("input:%v,wanted:%v,get:%v\n", v, outputs[k], res)
		}
	}
}

func Test_findUnsortedSubArray2(t *testing.T) {
	inputs := [][]int{
		{2, 6, 4, 8, 10, 9, 15},
	}
	outputs := []int{
		5,
	}

	for k, v := range inputs {
		res := findUnsortedSubarray2(v)
		if res != outputs[k] {
			t.Fatalf("input:%v,wanted:%v,get:%v\n", v, outputs[k], res)
		}
	}
}
