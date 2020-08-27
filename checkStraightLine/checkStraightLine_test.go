package checkStraightLine

import (
	"log"
	"testing"
)

func Test_checkStraightLine(t *testing.T) {
	inputs := [][][]int{
		{
			{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7},
		},
		{
			{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7},
		},
		{
			{0, 0}, {0, 1}, {0, -1},
		},
		{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}},
	}
	outputs := []bool{
		true, false, true, true,
	}

	for k, v := range inputs {
		tmp := checkStraightLine(v)
		if tmp != outputs[k] {
			log.Fatalf("input: %v, output: %v, wanted: %v\n", v, tmp, outputs[k])
		}
	}
}
