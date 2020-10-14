package arrayRankTransform

import (
	"log"
	"testing"
)

func Test_arrayRankTransform(t *testing.T) {
	inputs := [][]int{
		{40, 10, 20, 30},
		{100, 100, 100},
		{37, 12, 28, 9, 100, 56, 80, 5, 12},
	}
	outputs := [][]int{
		{4, 1, 2, 3},
		{1, 1, 1},
		{5, 3, 4, 2, 8, 6, 7, 1, 3},
	}

	for k, input := range inputs {
		res := arrayRankTransform(input)
		if !isEqual(res, outputs[k]) {
			log.Fatalf("\ninput: %v\nwanted: %v\nget: %v\n", input, outputs[k], res)
		}
	}
}

func isEqual(a, b []int) bool {
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
