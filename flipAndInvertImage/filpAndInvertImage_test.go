package flipAndInvertImage

import "testing"

func Test_flipAndInsertImage(t *testing.T) {
	test1 := [][]int{
		{1, 1, 1},
		{0, 0, 0},
		{1, 0, 1},
	}
	test1Required := [][]int{
		{0, 0, 0},
		{1, 1, 1},
		{0, 1, 0},
	}
	test2 := [][]int{
		{0, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	test2Required := [][]int{
		{0, 0, 1},
		{1, 0, 0},
		{0, 1, 0},
	}

	res1 := flipAndInvertImage(test1)
	if !_arrayEqual(test1Required, res1) {
		t.Error("test1 failed")
	}

	if !_arrayEqual(flipAndInvertImage(test2), test2Required) {
		t.Error("test2 failed")
	}
}

func _arrayEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		for kk, vv := range v {
			if b[k][kk] != vv {
				println("a:", vv, ",b:", b[k][kk])
				return false
			}
		}
	}

	return true
}
