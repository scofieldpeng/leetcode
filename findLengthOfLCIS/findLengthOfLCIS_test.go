package findLengthOfLCIS

import "testing"

func Test_findLengthOfLCIS(t *testing.T) {
	inputs := [][]int{
		{1, 2, 5, 4, 7},
		{2, 2, 2, 2, 2},
	}
	outouts := []int{
		3, 1,
	}

	for k, v := range inputs {
		res := findLengthOfLCIS(v)
		if res != outouts[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outouts[k], res)
		}
	}
}

func Test_findLengthOfLCIS2(t *testing.T) {
	inputs := [][]int{
		{1, 2, 5, 4, 7},
		{2, 2, 2, 2, 2},
	}
	outouts := []int{
		3, 1,
	}

	for k, v := range inputs {
		res := findLengthOfLCIS2(v)
		if res != outouts[k] {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outouts[k], res)
		}
	}
}
