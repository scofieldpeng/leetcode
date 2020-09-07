package shiftGrid

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_shiftGrid(t *testing.T) {
	args1 := [][][]int{
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{3, 8, 1, 9},
			{19, 7, 2, 5},
			{4, 6, 11, 10},
			{12, 0, 21, 13},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{1},
			{2},
			{3},
			{4},
			{7},
			{6},
			{5},
		},
	}
	args2 := []int{
		1, 4, 9, 23,
	}

	results := [][][]int{
		{
			{9, 1, 2},
			{3, 4, 5},
			{6, 7, 8},
		},
		{
			{12, 0, 21, 13},
			{3, 8, 1, 9},
			{19, 7, 2, 5},
			{4, 6, 11, 10},
		},
		{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		{
			{6},
			{5},
			{1},
			{2},
			{3},
			{4},
			{7},
		},
	}

	for k := range args1 {
		result := shiftGrid(args1[k], args2[k])
		hash1, _ := json.Marshal(result)
		hash2, _ := json.Marshal(results[k])
		if fmt.Sprintf("%x", hash1) != fmt.Sprintf("%x", hash2) {
			t.Fatalf("shift failed\nsource:%v\nk:%d\nwanted: %v\nget: %v\n", args1[k], args2[k], results[k], result)
		}
	}
}
