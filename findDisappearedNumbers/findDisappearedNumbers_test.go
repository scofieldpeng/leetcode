package findDisappearedNumbers

import "testing"

func Test_findDisappearedNumbers(t *testing.T) {
	inputs := [][]int{
		{4, 3, 2, 7, 8, 2, 3, 1},
	}
	outputs := [][]int{
		{5, 6},
	}

	for k, v := range inputs {
		res := findDisappearedNumbers(v)
		if false == equal(res, outputs[k]) {
			t.Fatalf("input:%v, output:%v, wanted:%v\n", v, res, outputs[k])
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}
