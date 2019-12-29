package isOneBitCharacter

import "testing"

func Test_isOneBitCharacter(t *testing.T) {
	inputs := [][]int{
		{1, 0, 0},
		{1, 1, 1, 0},
	}
	outputs := []bool{
		true, false,
	}

	for k, v := range inputs {
		res := isOneBitCharacter(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}

func Test_isOneBitCharacter2(t *testing.T) {
	inputs := [][]int{
		{1, 0, 0},
		{1, 1, 1, 0},
	}
	outputs := []bool{
		true, false,
	}

	for k, v := range inputs {
		res := isOneBitCharacter2(v)
		if res != outputs[k] {
			t.Fatalf("input: %v, wanted: %v, get: %v\n", v, outputs[k], res)
		}
	}
}
