// problem: https://leetcode.com/problems/missing-number/
// solution: https://algs.home.pjf.im/missing-number.html
package missingNumber

import "sort"

// sort
func missingNumber(nums []int) int {
	sort.Ints(nums)
	prev := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if prev+1 != nums[i] {
			return prev + 1
		}
		prev = nums[i]
	}

	return -1
}

// space change time
func missingNumber2(nums []int) int {
	length := len(nums)
	numsArr := make([]int, length+1, length+1)

	for i := 0; i < length; i++ {
		numsArr[nums[i]] = nums[i]
	}

	for k, v := range numsArr {
		if v == 0 && k != 0 {
			return k
		}
	}

	return -1
}
