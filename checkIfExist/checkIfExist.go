// problem: https://leetcode.com/problems/check-if-n-and-its-double-exist/
package checkIfExist

import "sort"

func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == arr[i]*2 || arr[i] == arr[j]*2 {
				return true
			}
		}
	}

	return false
}

func checkIfExist2(arr []int) bool {
	arrMap := make(map[int]bool)
	for _, v := range arr {
		if _, ok := arrMap[v*2]; ok {
			return true
		}
		if v%2 == 0 {
			if _, ok := arrMap[v/2]; ok {
				return true
			}
		}
		arrMap[v] = true
	}

	return false
}
