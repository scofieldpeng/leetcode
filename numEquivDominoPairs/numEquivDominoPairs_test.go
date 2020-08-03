package numEquivDominoPairs

import "testing"

func Test_numEquivDominoPairs(t *testing.T) {
	inputs := [][][]int{
		{{1, 2}, {2, 1}, {3, 4}, {5, 6}},
		{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}},
		{{1, 1}, {2, 2}, {1, 1}, {1, 2}, {1, 2}, {1, 1}},
	}

	outputs := []int{
		1,
		3,
		4,
	}

	for k, v := range inputs {
		res := numEquivDominoPairs(v)
		if res != outputs[k] {
			t.Fatalf("input:%v \n output:%v \n wanted:%v", v, res, outputs[k])
		}
	}
}
