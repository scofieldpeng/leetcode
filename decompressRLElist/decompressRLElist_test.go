// problem: https://leetcode.com/problems/decompress-run-length-encoded-list/
package decompressRLElist

import "testing"

func Test_decompressRLElist(t *testing.T) {
	inputs := [][]int{
		{1, 2, 3, 4},
		{1, 1, 2, 3},
	}
	outputs := [][]int{
		{2, 4, 4, 4},
		{1, 3, 3},
	}

	for k, input := range inputs {
		res := decompressRLElist(input)
		if !isEqual(res, outputs[k]) {
			t.Fatalf("\ninput: %v\nget: %v\nwanted: %v\n", input, res, outputs[k])
		}
	}
}

func isEqual(a, b []int) bool {
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
