// source: https://leetcode.com/problems/search-insert-position/
// solution:
package searchInsert

import "fmt"

func searchInsert(nums []int, target int) int {
	length := len(nums)
	// 首
	if nums[0] >= target {
		return 0
	}
	// 尾
	if nums[length-1] < target {
		return length
	}

	var (
		start, end = 0, length - 1
		mid        = 0
	)

	for ; start <= end; {
		// equal to (end - start) / 2
		mid = (end + start) >> 1

		fmt.Printf("start:%v(%v), mid:%v(%v), target:%v, end:%v(%v)\n", start, nums[start], mid, nums[mid], target, end, nums[end])

		if nums[mid] == target {
			return mid
		}

		// 这里+1，-1是为了防止当只有两个元素的时候，mid永远是左边的那个值
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return start
}
