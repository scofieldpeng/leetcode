package twoSumII

import "testing"

func Test_twoSum(t *testing.T) {
	arg1 := [][]int{
		{2, 7, 11, 15},
	}
	args := []int{
		9,
	}
	output := [][]int{
		{1, 2},
	}

	for k, v := range arg1 {
		res := twoSum(v, args[k])
		if !equal(res, output[k]) {
			t.Fatalf("input:%v,%v, output:%v, wanted:%v\n", arg1[k], args[k], output[k], res)
		}
	}
}

func equal(data1, data2 []int) bool {
	if len(data1) != len(data2) {
		return false
	}

	for k, v := range data1 {
		if v != data2[k] {
			return false
		}
	}

	return true
}
