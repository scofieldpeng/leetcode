package arrayPairSum

import "testing"

func Test_arrayPairSum(t *testing.T) {
	input := []int{1, 4, 3, 2}
	output := 4

	if arrayPairSum(input) != output {
		t.Error("error,want:", output, ",get:", arrayPairSum(input))
	}
}