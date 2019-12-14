package canPlaceFlowers

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	input1 := [][]int{
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}
	input2 := []int{
		1, 2,
	}

	outputs := []bool{
		true, false,
	}

	for k, v := range input1 {
		res := canPlaceFlowers(v, input2[k])
		if res != outputs[k] {
			t.Fatalf("input: %v,%v, wanted: %v, get: %v\n", v, input2[k], outputs[k], res)
		}
	}
}
