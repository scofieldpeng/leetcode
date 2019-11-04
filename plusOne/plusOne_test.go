package plusOne

import "testing"

func Test_plusOne(t *testing.T) {
	asserts := [][][]int{
		{
			{1, 2, 3},
			{1, 2, 4},
		},
		{
			{1, 1, 9},
			{1, 2, 0},
		},
		{
			{1, 9, 9},
			{2, 0, 0},
		},
	}
	for _, v := range asserts {
		res := plusOne(v[0])

		if !arrayEqual(res, v[1]) {
			t.Fatalf("input: %v, want: %v, get:%v\n", v[0], v[1], res)
		}
	}
}

func arrayEqual(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for k, v := range arr1 {
		if v != arr2[k] {
			return false
		}
	}

	return true
}
