package moveZeroes

import "testing"

func Test_moveZeros(t *testing.T) {
	inputs := [][]int{
		{0, 1, 0, 3, 12},
	}
	output := [][]int{
		{1, 3, 12, 0, 0},
	}

	for k, v := range inputs {
		moveZeroes(v)
		if false == equals(v, output[k]) {
			t.Fatalf("wanted: %v, get: %v", output[k], v)
		}
	}
}

func equals(a, b []int) bool {
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
