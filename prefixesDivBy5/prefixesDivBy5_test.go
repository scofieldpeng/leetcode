package prefixesDivBy5

import "testing"

func Test_prefixesDivBy5(t *testing.T) {
	inputs := [][]int{
		// {0, 1, 1},
		// {1, 1, 1},
		// {0, 1, 1, 1, 1, 1},
		// {1, 1, 1, 0, 1},
		// {0,1,0,1,0,0,1,0,0,1,0,1,0,0,1,1,0,1,1,1,0,0,0,0,0,1,0,1,1},
		{1,0,1,1,1,1,0,0,0,0,1,0,0,0,0,0,1,0,0,1,1,1,1,1,0,0,0,0,1,1,1,0,0,0,0,0,1,0,0,0,1,0,0,1,1,1,1,1,1,0,1,1,0,1,0,0,0,0,0,0,1,0,1,1,1,0,0,1,0},
	}
	outputs := [][]bool{
		// {true, false, false},
		// {false, false, false},
		// {true, false, false, false, true, false},
		// {false, false, false, false, false},
		// {true,false,false,true,true,true,false,false,false,false,false,false,false,false,false,true,true,false,false,false,false,false,false,false,false,false,false,false,true},
		{false,false,true,false,false,false,false,false,false,false,true,true,true,true,true,true,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,false,true,false,false,false,true,false,false,true,false,false,true,true,true,true,true,true,true,false,false,true,false,false,false,false,true,true},
	}

	for k, v := range inputs {
		res := prefixesDivBy5(v)
		if false == equals(res, outputs[k]) {
			t.Fatalf("\ninput:%v\nget:%v\nwanted:%v\n", v, res, outputs[k])
		}
	}
}

func equals(a, b []bool) bool {
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
