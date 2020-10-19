package checkIfExist

import "testing"

func Test_checkIfExist(t *testing.T) {
	inputs := [][]int{
		{10, 2, 5, 3},
		{7, 1, 14, 11},
		{3, 1, 7, 11},
		{-10, 12, -20, -8, 15},
	}
	outputs := []bool{
		true,
		true,
		false,
		true,
	}

	for k, v := range inputs {
		res := checkIfExist(v)
		if res != outputs[k] {
			t.Fatalf("input: %v\nwanted:%v\nget:%v\n", v, outputs[k], res)
		}
	}
}
func Test_checkIfExist2(t *testing.T) {
	inputs := [][]int{
		{10, 2, 5, 3},
		{7, 1, 14, 11},
		{3, 1, 7, 11},
		{-10, 12, -20, -8, 15},
	}
	outputs := []bool{
		true,
		true,
		false,
		true,
	}

	for k, v := range inputs {
		res := checkIfExist2(v)
		if res != outputs[k] {
			t.Fatalf("input: %v\nwanted:%v\nget:%v\n", v, outputs[k], res)
		}
	}
}
