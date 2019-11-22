// problem: https://leetcode.com/problems/missing-number/
// solution: https://algs.home.pjf.im/missing-number.html
package missingNumber

import "sort"

// sort
func missingNumber(nums []int) int {
	sort.Ints(nums)
	// 第一个不存在
	if nums[0] != 0 {
		return 0
	}
	// 最后一个不存在
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}

	// 中间有值不存在
	prev := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if prev+1 != nums[i] {
			return prev + 1
		}
		prev = nums[i]
	}

	// 都存在返回-1
	return -1
}

// space change time
func missingNumber2(nums []int) int {
	length := len(nums)
	numsArr := make([]int, length+1, length+1)

	findZero := false

	for i := 0; i < length; i++ {
		if nums[i] == 0 {
			findZero = true
		}
		numsArr[nums[i]] = nums[i]
	}

	// 看是否存在0，不存在直接返回
	if !findZero {
		return 0
	}

	for k, v := range numsArr {
		if v == 0 && k != 0 {
			return k
		}
	}

	return -1
}

// 利用高斯算法
func missingNumber3(nums []int) int {
	length := len(nums)
	total := length * (length + 1) / 2
	for _, v := range nums {
		total -= v
	}

	return total
}

// 利用位运算的异或来解决
// a ^ a = 0
// a ^ 0 = a
// a ^ b ^ a = b
func missingNumber4(nums []int) int {
	lastV := len(nums)
	for k, v := range nums {
		lastV = lastV ^ k ^ v
	}

	return lastV
}