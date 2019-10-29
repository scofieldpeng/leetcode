package duplicateZeros

import "testing"

func Test_duplicateZeroNumbers(t *testing.T) {
	arr1 := []int{2, 1, 0, 3}
	arr2 := []int{2, 1, 0}
	arr3 := []int{0, 0, 0}
	arr4 := []int{1, 2, 3}

	res1,_ := duplicateZeroNumbers(arr1)
	if res1 != 1 {
		t.Fatalf("2,1,0,3 calculate error,want 1,get %d\n", res1)
	}
	t.Logf("2,1,0,3 get %d\n", res1)

	res2,_ := duplicateZeroNumbers(arr2)
	if res2 != 0 {
		t.Fatalf("2,1,0 calculate error,want 0,get %d\n", res2)
	}
	t.Logf("2,1,0 get %d\n", res2)

	res3,_ := duplicateZeroNumbers(arr3)
	if res3 != 2 {
		t.Fatalf("0,0,0 calculate error,want 2,get %d\n", res3)
	}
	t.Logf("0,0,0 get %d\n", res2)

	res4,_ := duplicateZeroNumbers(arr4)
	if res4 != 0 {
		t.Fatalf("0,0,0 calculate error,want 0,get %d\n", res4)
	}
	t.Logf("1,2,3 get %d\n", res4)
}

func Test_duplicateZeros(t *testing.T) {
	arr1 := []int{2, 1, 0, 3}
	arr2 := []int{2, 1, 0}
	arr3 := []int{0, 0, 0}
	arr4 := []int{1, 2, 3}
	arr5 := []int{0, 1, 7, 6, 0, 2, 0, 7}
	arr6 := []int{8, 4, 5, 0, 0, 0, 0, 7}
	arr7 := []int{9, 9, 9, 4, 8, 0, 0, 3, 7, 2, 0, 0, 0, 0, 9, 1, 0, 0, 1, 1, 0, 5, 6, 3, 1, 6, 0, 0, 2, 3, 4, 7, 0, 3, 9, 3, 6, 5, 8, 9, 1, 1, 3, 2, 0, 0, 7, 3, 3, 0, 5, 7, 0, 8, 1, 9, 6, 3, 0, 8, 8, 8, 8, 0, 0, 5, 0, 0, 0, 3, 7, 7, 7, 7, 5, 1, 0, 0, 8, 0, 0}
	arr8 := []int{1,5,2,0,6,8,0,6,0}

	duplicateZeros(arr1)
	if !compareSlice(arr1, []int{2, 1, 0, 0}) {
		t.Fatalf("[2,1,0,3] error, get:%v", arr1)
	}
	t.Logf("[2,1,0,3] get:%v", arr1)

	duplicateZeros(arr2)
	if !compareSlice(arr2, []int{2, 1, 0,}) {
		t.Fatalf("[2,1,0] error, get:%v", arr2)
	}
	t.Logf("[2,1,0] get:%v", arr2)

	duplicateZeros(arr3)
	if !compareSlice(arr3, []int{0, 0, 0}) {
		t.Fatalf("[0,0,0] error, get:%v", arr3)
	}
	t.Logf("[0,0,0] get:%v", arr3)

	duplicateZeros(arr4)
	if !compareSlice(arr4, []int{1, 2, 3}) {
		t.Fatalf("[1,2,3] error, get:%v", arr4)
	}
	t.Logf("[1,2,3] get:%v", arr4)

	duplicateZeros(arr5)
	if !compareSlice(arr5, []int{0, 0, 1, 7, 6, 0, 0, 2}) {
		t.Fatalf("[0,1,7,6,0,2,0,7] error, get:%v", arr5)
	}
	t.Logf("[0,1,7,6,0,2,0,7] get:%v", arr5)

	duplicateZeros(arr6)
	if !compareSlice(arr6, []int{8, 4, 5, 0, 0, 0, 0, 0}) {
		t.Fatalf("[8,4,5,0,0,0,0,7] error, get:%v", arr6)
	}
	t.Logf("[8,4,5,0,0,0,0,7] get:%v", arr6)

	duplicateZeros(arr7)
	if !compareSlice(arr7, []int{9, 9, 9, 4, 8, 0, 0, 0, 0, 3, 7, 2, 0, 0, 0, 0, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 1, 1, 0, 0, 5, 6, 3, 1, 6, 0, 0, 0, 0, 2, 3, 4, 7, 0, 0, 3, 9, 3, 6, 5, 8, 9, 1, 1, 3, 2, 0, 0, 0, 0, 7, 3, 3, 0, 0, 5, 7, 0, 0, 8, 1, 9, 6, 3, 0, 0, 8, 8, 8, 8, 0}) {
		t.Fatalf("[9,9,9,4,8,0,0,3,7,2,0,0,0,0,9,1,0,0,1,1,0,5,6,3,1,6,0,0,2,3,4,7,0,3,9,3,6,5,8,9,1,1,3,2,0,0,7,3,3,0,5,7,0,8,1,9,6,3,0,8,8,8,8,0,0,5,0,0,0,3,7,7,7,7,5,1,0,0,8,0,0] error, get:%v", arr7)
	}
	t.Logf("[9,9,9,4,8,0,0,3,7,2,0,0,0,0,9,1,0,0,1,1,0,5,6,3,1,6,0,0,2,3,4,7,0,3,9,3,6,5,8,9,1,1,3,2,0,0,7,3,3,0,5,7,0,8,1,9,6,3,0,8,8,8,8,0,0,5,0,0,0,3,7,7,7,7,5,1,0,0,8,0,0] get:%v", arr7)

	duplicateZeros(arr8)
	if !compareSlice(arr8, []int{1,5,2,0,0,6,8,0,0}) {
		t.Fatalf("[1,5,2,0,6,8,0,6,0] error, get:%v", arr8)
	}
	t.Logf("[1,5,2,0,6,8,0,6,0] get:%v", arr8)
}

func compareSlice(sliceA, sliceB []int) bool {
	if len(sliceA) != len(sliceB) {
		return false
	}

	for k, v := range sliceA {
		if v != sliceB[k] {
			return false
		}
	}

	return true
}
