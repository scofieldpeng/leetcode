package smallerNumbersThanCurrent

import (
	"log"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	inputs := [][]int{
		{8, 1, 2, 2, 3},
		{6, 5, 4, 8},
		{7, 7, 7, 7},
	}
	outputs := [][]int{
		{4, 0, 1, 1, 3},
		{2, 1, 0, 3},
		{0, 0, 0, 0},
	}

	for k, input := range inputs {
		res := smallerNumbersThanCurrent(input)
		if !isEqual(res, outputs[k]) {
			log.Fatalf("\ninput:%v\noutput:%v\nwanted:%v\n", input, res, outputs[k])
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
