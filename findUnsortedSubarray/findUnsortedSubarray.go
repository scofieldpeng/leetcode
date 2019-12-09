package findUnsortedSubarray

import "sort"

// sort大法
func findUnsortedSubarray(nums []int) int {
	var data = make([]int, len(nums), len(nums))
	copy(data, nums)
	sort.Ints(data)

	begin := 0
	end := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != data[i] {
			begin = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != data[i] {
			end = i
			break
		}
	}

	return end - begin + 1
}
