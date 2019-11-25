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
	for _, v := range nums {
		if v != 0 {
			nums[lastNoZeroIndex] = v
			lastNoZeroIndex++
		}
	}

	for i := lastNoZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}
