package main

import (
	"fmt"
	"testing"
)

func Test_relativeSortArray(t *testing.T) {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19,}
	arr2 := []int{2, 1, 4, 3, 9, 6,}
	// expect := []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19,}

	res := relativeSortArray(arr1, arr2)
	fmt.Println(res)
}
