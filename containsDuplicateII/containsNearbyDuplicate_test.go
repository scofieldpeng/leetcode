package containsDuplicateII

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	input1 := [][]int{
		{1, 2, 3, 1},
		{1, 0, 1, 1},
		{1, 2, 3, 1, 2, 3},
	}
	input2 := []int{3, 1, 2,}
	outputs := []bool{true, true, false}

	for k, v := range input1 {
		res := containsNearbyDuplicate(v, input2[k])
		if res != outputs[k] {
			t.Fatalf("input: %v,%v, wanted: %v,get: %v\n", input1[k], input2[k], outputs[k], res)
		}
	}
}
