// problem: https://leetcode.com/problems/move-zeroes/
// solution: https://algs.home.pjf.im/move-zeroes.html
package moveZeroes

// 简单粗暴，用一个数组来存储0的索引值，然后将非空的值移动到这个索引位即可，移走的那位需要更新成0
func moveZeroes(nums []int) {
	zeroIndexs := make([]int, 0, len(nums))
	for k, v := range nums {
		if v == 0 {
			zeroIndexs = append(zeroIndexs, k)
			continue
		}

		if len(zeroIndexs) > 0 {
			nums[zeroIndexs[0]] = v
			zeroIndexs = zeroIndexs[1:]
			zeroIndexs = append(zeroIndexs, k)
			nums[k] = 0
		}
	}
}

func moveZeros2(nums []int) {
	lastNoZeroIndex := 0
	// 将非0的值向前面挪动
	for _, v := range nums {
		// 这里非常精妙，挪动非0值后，将其指向下一位，如果下一位是0，那么会停住不动，直到遍历到末尾或者是遇到了非0的值
		if v != 0 {
			nums[lastNoZeroIndex] = v
			lastNoZeroIndex++
		}
	}
	// 从最后一个非0值的下一位开始，将剩余的所有下标的值都变为0
	for i := lastNoZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeros3(nums []int) {
	lastNoZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		// 当前位置不为0的时候，考虑将当前值挪动到找到的第一个0值位置去
		if nums[i] != 0 {
			if lastNoZeroIndex != i {
				nums[lastNoZeroIndex] = nums[i]
				nums[i] = 0
			}
			lastNoZeroIndex++
		}
	}
}
