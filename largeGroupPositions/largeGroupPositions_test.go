package largeGroupPositions

import "testing"

func Test_largeGroupPositons(t *testing.T) {
	inputs := []string{
		"abbxxxxzzy",
		"abc",
		"abcdddeeeeaabbbcd",
		"aaa",
	}

	outputs := [][][]int{
		{{3, 6}},
		{},
		{{3, 5}, {6, 9}, {12, 14}},
		{{0, 2}},
	}

	for k, v := range inputs {
		res := largeGroupPositions(v)
		if false == equal(res, outputs[k]) {
			t.Fatalf("input:%v, wanted:%v, get:%v\n", v, outputs[k], res)
		}
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if len(v) != len(b[k]) {
			return false
		}

		for kk, vv := range v {
			if vv != b[k][kk] {
				return false
			}
		}
	}

	return true
}
