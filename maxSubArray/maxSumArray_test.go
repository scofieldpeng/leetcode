package maxSubArray

import "testing"

func Test_maxSubArray(t *testing.T) {
	asserts := [][]int{
		{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		{6},
	}

	for i := 0; i < len(asserts); i += 2 {
		res := maxSubArray(asserts[i])

		if res != asserts[i+1][0] {
			t.Fatalf("input: %v, want: %v, get: %v", asserts[i], asserts[i+1][0], res)
		}
	}
}
