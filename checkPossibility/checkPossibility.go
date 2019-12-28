package checkPossibility

// 遍历的时候找到所有开始decr的值，然后看这个值能否改为一个新值，这个新值是否大于上一个值，并且小于下一个值，如果没有，直接返回false
// 如果有那么更新到新值继续遍历
// 如果在找到之后还发现了有新的decr，那么直接false
func checkPossibility(nums []int) bool {
	length := len(nums)
	wrongIndex := -1
	// 先看出现了多少个错位的值
	for current := 1; current < length; current++ {
		if nums[current-1] > nums[current] {
			if wrongIndex == -1 {
				wrongIndex = current
				continue
			}
			return false
		}
	}

	// 全部都是按顺序排列
	// 如果第二位不按顺序，改为可以和前一位或者后一位相同
	// 如果是最后一位不按顺序，改为和倒数第二位一致即可
	if wrongIndex == -1 || wrongIndex == 1 || wrongIndex == length-1 {
		return true
	}

	// 对于[-1,4,1,2]这种，需要改变的是4
	if nums[wrongIndex-2] <= nums[wrongIndex] {
		return true
	}
	// 对于[3,4,2,6]这种，需要改变的是2
	if nums[wrongIndex+1] >= nums[wrongIndex-1] {
		return true
	}

	return false
}
