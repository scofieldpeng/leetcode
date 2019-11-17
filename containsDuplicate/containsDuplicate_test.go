package containsDuplicate

import (
	"log"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 2, 3, 4},
		{3, 2, 5, 1, 1, 7},
		{2, 5, 1, 9, 8, 8},
	}
	outputs := []bool{
		false, true, true, true,
	}

	for k, v := range inputs {
		res := containsDuplicate(v)
		if res != outputs[k] {
			log.Fatalf("input: %v, output: %v, wanted: %v\n", v, res, outputs[k])
		}
	}
}

func Test_containsDuplicate2(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 2, 3, 4},
		{3, 2, 5, 1, 1, 7},
		{2, 5, 1, 9, 8, 8},
	}
	outputs := []bool{
		false, true, true, true,
	}

	for k, v := range inputs {
		res := containsDuplicate2(v)
		if res != outputs[k] {
			log.Fatalf("input: %v, output: %v, wanted: %v\n", v, res, outputs[k])
		}
	}
}
