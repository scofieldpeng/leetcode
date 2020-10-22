package countNegatives

import (
	"log"
	"testing"
)

func Test_countNegatives(t *testing.T) {
	inputs := [][][]int{
		{
			{4, 3, 2, -1},
			{3, 2, 1, -1},
			{1, 1, -1, -2},
			{-1, -1, -2, -3},
		},
		{
			{3, 2},
			{1, 0},
		},
		{
			{1, -1},
			{-1, -1},
		},
		{
			{-1},
		},
	}
	outputs := []int{
		8, 0, 3, 1,
	}

	for k, input := range inputs {
		res := countNegatives(input)
		if res != outputs[k] {
			log.Fatalf("\ninput: %v\noutput: %v\nwanted: %v\n", input, res, outputs[k])
		}
	}
}
