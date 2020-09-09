package minTimeToVisitAllPoints

import (
	"log"
	"testing"
)

func Test_minTimeToVisitAllPoints(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 1},
			{3, 4},
			{-1, 0},
		},
		{
			{3, 2},
			{-2, 2},
		},
	}
	outputs := []int{
		7, 5,
	}

	for k, input := range inputs {
		res := minTimeToVisitAllPoints(input)
		if res != outputs[k] {
			log.Fatalf("input: %v\nwanted: %v\nget: %v\n", input, outputs[k], res)
		}
	}
}
