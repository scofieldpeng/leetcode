// problem: https://leetcode.com/proble1ms/the-k-weakest-rows-in-a-matrix/
package kWeakestRows

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	// 获取到每行士兵的数量
	soldierArr := make([]int, len(mat), len(mat))
	soldierMap := make(map[int][]int)
	for k, v := range mat {
		num := 0
		for _, vv := range v {
			if vv == 0 {
				break
			}
			num++
		}
		soldierArr[k] = num
		if _, ok := soldierMap[num]; !ok {
			soldierMap[num] = make([]int, 0)
		}
		soldierMap[num] = append(soldierMap[num], k)
	}

	// 按照士兵数量排序
	soldierSortArr := make([]int, len(mat), len(mat))
	for k, v := range soldierArr {
		soldierSortArr[k] = v
	}
	sort.Ints(soldierSortArr)

	// 二维拍平
	tmp := make([]int, 0)
	for i := 0; i < len(soldierSortArr); {
		tmp2 := soldierMap[soldierSortArr[i]]
		tmp = append(tmp, tmp2...)
		i += len(tmp2)
	}

	return tmp[0:k]
}
func kWeakestRows2(mat [][]int, k int) []int {
	windex := make([]int, len(mat))
	n := len(mat[0])
	for i := 0; i < len(mat); i++ {
		weight := 0
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				break
			}
			weight++
		}
		// 先左移16位( 等同weight * (2^16) )是为了让数组尽可能大，然后一个或运算，后可以让有相同士兵的行按照索引值排序
		windex[i] = weight<<16 | i
	}
	sort.Ints(windex)
	for i := 0; i < k; i++ {
		// 这里一个&运算符反解析出i的值
		windex[i] &= 0xffff
	}
	return windex[:k]
}
