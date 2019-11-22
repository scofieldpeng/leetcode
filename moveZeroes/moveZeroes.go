// problem: https://leetcode.com/problems/move-zeroes/
// solution: https://algs.home.pjf.im/move-zeroes.html
package moveZeroes

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
