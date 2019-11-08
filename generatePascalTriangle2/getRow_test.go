package generatePascalTriangle2

import (
	"testing"
)

func Test_getRow(t *testing.T) {
	inputs := []int{0, 1, 2, 3, 4,}
	outputs := [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}
	for k, v := range inputs {
		output := getRow(v)
		if false == equal(output, outputs[k]) {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], output)
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
