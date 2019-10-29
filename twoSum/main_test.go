package main

import "testing"

func Test_TwoSum(t *testing.T) {
	v := []int{1, 3, 5, 7}
	expect := 6
	if res := twoSum(v, expect); len(res) == 2 && v[res[0]]+v[res[1]] == expect {
		println("ok")
	} else {
		t.Error("should find,but not found")
	}
	expect = 15
	if res := twoSum(v, expect); len(res) == 0 {
		println("ok")
	} else {
		t.Error("should not found, but find")
	}
}
