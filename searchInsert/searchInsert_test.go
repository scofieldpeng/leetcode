package searchInsert

import "testing"

func Test_searchInsert(t *testing.T) {
	asserts := [][]int{
		{1, 3, 5, 6},
		{5, 2},
		{1, 3, 5, 6},
		{2, 1},
		{1, 3, 5, 6},
		{7, 4},
		{1, 3, 5, 6},
		{0, 0},
		{1,3},
		{3,1},
	}

	for i := 0; i < len(asserts); i += 2 {
		res := searchInsert(asserts[i], asserts[i+1][0])
		if res != asserts[i+1][1] {
			t.Fatalf("source: %v, arg: %v, wanted: %v, get: %v\n", asserts[i], asserts[i+1][0], asserts[i+1][1], res)
		}
	}
}
