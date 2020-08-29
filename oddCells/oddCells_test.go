package oddCells

import (
	"log"
	"testing"
)

func Test_oddCells(t *testing.T) {
	inputs1 := [][][]int{
		{
			{0, 1},
			{1, 1},
		},
		{
			{1, 1},
			{0, 0},
		},
	}
	inputs2 := [][]int{
		{2, 3},
		{2, 2},
	}

	outputs := []int{
		6, 0,
	}
	for i, input1 := range inputs1 {
		res := oddCells(inputs2[i][0], inputs2[i][1], input1)
		if res != outputs[i] {
			log.Fatalf("inputs: %v,n: %v, m: %v, wanted: %v, get: %v \n", input1, inputs2[i][0], inputs2[i][1], outputs[i], res)
		}
	}
}

func Test_oddCells2(t *testing.T) {
	inputs1 := [][][]int{
		{
			{0, 1},
			{1, 1},
		},
		{
			{1, 1},
			{0, 0},
		},
	}
	inputs2 := [][]int{
		{2, 3},
		{2, 2},
	}

	outputs := []int{
		6, 0,
	}
	for i, input1 := range inputs1 {
		res := oddCells2(inputs2[i][0], inputs2[i][1], input1)
		if res != outputs[i] {
			log.Fatalf("inputs: %v,n: %v, m: %v, wanted: %v, get: %v \n", input1, inputs2[i][0], inputs2[i][1], outputs[i], res)
		}
	}
}
