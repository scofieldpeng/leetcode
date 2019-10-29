package matrixReshape

import (
	"fmt"
	"testing"
)

func Test_matrixReshape(t *testing.T) {
	nums := [][]int{{1, 2}, {3, 4}}
	r := 4
	c := 1

	res := matrixReshape(nums, r, c)
	fmt.Println(res)
}
