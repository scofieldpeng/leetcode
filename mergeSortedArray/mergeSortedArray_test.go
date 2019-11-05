package mergeSortedArray

import (
	"fmt"
	"testing"
)

func Test_mergeSortedArray(t *testing.T) {
	asserts := [][]int{
		{1, 2, 3, 0, 0, 0},
		{2, 5, 6},
		{3, 3},
		{1, 2, 2, 3, 5, 6},
		{4, 5, 6, 0, 0, 0},
		{1, 2, 3},
		{3, 3},
		{1, 2, 3, 4, 5, 6},
	}

	for i := 0; i < len(asserts); i += 4 {
		nums1 := asserts[i]
		merge(asserts[i], asserts[i+2][0], asserts[i+1], asserts[i+2][1])
		fmt.Println(nums1)

		if !equal(asserts[i], nums1) {
			t.Fatalf("nums1: %v, nums2: %v, m: %v, n: %v, wanted: %v, get: %v\n", asserts[i], asserts[i+1], asserts[i+2][0], asserts[i+2][1], asserts[i+3], nums1)
		}
	}
}

func equal(data1, data2 []int) bool {
	if len(data1) != len(data2) {
		return false
	}

	for k, v := range data1 {
		if v != data2[k] {
			return false
		}
	}

	return true
}
