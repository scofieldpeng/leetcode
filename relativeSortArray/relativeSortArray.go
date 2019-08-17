package main

import (
	"sort"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	var tmp [1001]int

	var leftArr = make([]int, 0, len(arr1)-len(arr2))
	for _, v := range arr1 {
		// 没有在arr2的值
		if !inArray(v, arr2) {
			leftArr = append(leftArr, v)
			continue
		}

		tmp[v]++
	}

	if len(leftArr) > 0 {
		sort.Ints(leftArr)
	}

	res := make([]int, 0, len(arr1)-len(arr2))
	for _, v := range arr2 {
		if tmp[v] > 0 {
			for j := 1; j <= tmp[v]; j++ {
				res = append(res, v)
			}
		}
	}

	return append(res, leftArr...)
}

func inArray(find int, arr []int) bool {
	for _, v := range arr {
		if v == find {
			return true
		}
	}

	return false
}

func relativeSortArray2(arr1 []int, arr2 []int) []int {
	max := arr1[0]
	// 这里先确定arr1中最大的值是多少，这个值后期是方便接下俩申请对应长度的slice，避免占用太大的内存
	for _, value := range arr1 {
		if value > max {
			max = value
		}
	}

	// 申请对应的长度的数组
	temp := make([]int, max+1)
	// 这里遍历arr1，将值记录到对应的下标上，如果存在多个，那么就将其的值+1
	for _, value := range arr1 {
		temp[value]++
	}

	res := []int{}
	for _, val := range arr2 {
		num := temp[val]
		// 按照arr2 输出结果
		for i := 0; i < num; i++ {
			res = append(res, val)
		}
		// 这里要记得清空为0，为接下来的arr1和arr2的差集做准备
		temp[val] = 0
	}

	for key, val := range temp {
		if val != 0 {
			for i := 0; i < val; i++ {
				res = append(res, key)
			}
		}
	}

	return res
}
