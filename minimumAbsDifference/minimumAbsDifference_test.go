package minimumAbsDifference

import (
	"encoding/json"
	"testing"
)

func Test_minimumAbsDifference(t *testing.T) {
	inputs := [][]int{
		{4, 2, 1, 3},
		{1, 3, 6, 10, 15},
		{3, 8, -10, 23, 19, -4, -14, 27},
	}
	outputs := [][][]int{
		{
			{1, 2},
			{2, 3},
			{3, 4},
		},
		{
			{1, 3},
		},
		{
			{-14, -10},
			{19, 23},
			{23, 27},
		},
	}

	for k, v := range inputs {
		output := minimumAbsDifference(v)
		if len(output) != len(outputs[k]) {
			t.Fatalf("failed! k:%v,v:%v,wanted:%v,get:%v\n", k, v, outputs[k], output)
		}
		get, _ := json.Marshal(output)
		wanted, _ := json.Marshal(outputs[k])
		if string(get) != string(wanted) {
			t.Fatalf("failed2! k:%v,v:%v,wanted:%v,get:%v\n", k, v, outputs[k], output)
		}
	}
}
