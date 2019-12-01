package findPairs

import "testing"

func Test_findPairs(t *testing.T) {
	arg1 := [][]int{
		{3, 1, 4, 1, 5},
		{1, 2, 3, 4, 5},
		{1, 3, 1, 5, 4},
		{1, 1, 1, 2, 1},
	}
	arg2 := []int{
		2, 1, 0, 1,
	}

	outputs := []int{
		2, 4, 1, 1,
	}

	for k, v := range arg1 {
		res := findPairs(v, arg2[k])
		if res != outputs[k] {
			t.Fatalf("input:%v,%v,wanted:%v,get:%v\n", v, arg2[k], outputs[k], res)
		}
	}
}

func Test_findPairs2(t *testing.T) {
	arg1 := [][]int{
		{3, 1, 4, 1, 5},
		{1, 2, 3, 4, 5},
		{1, 3, 1, 5, 4},
		{1, 1, 1, 2, 1},
	}
	arg2 := []int{
		2, 1, 0, 1,
	}

	outputs := []int{
		2, 4, 1, 1,
	}

	for k, v := range arg1 {
		res := findPairs2(v, arg2[k])
		if res != outputs[k] {
			t.Fatalf("input:%v,%v,wanted:%v,get:%v\n", v, arg2[k], outputs[k], res)
		}
	}
}
