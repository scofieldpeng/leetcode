// problem: https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
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

// todo 待分析
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
		// 这里是
		windex[i] = weight<<16 | i
	}
	sort.Ints(windex)
	for i := 0; i < k; i++ {
		windex[i] &= 0xffff
	}
	return windex[:k]
}
