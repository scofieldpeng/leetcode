// problem: https://leetcode.com/problems/largest-number-at-least-twice-of-others/
// solution: https://algs.home.pjf.im/largest-number-at-least-twice-of-others.html
package dominantIndex

func dominantIndex(nums []int) int {
	// find the max
	maxIndex := -1
	for k, v := range nums {
		if v >= nums[maxIndex] {
			maxIndex = k
			continue
		}
	}
	if maxIndex > -1 {
		for k, v := range nums {
			if maxIndex == k {
				continue
			}

			if v*2 > nums[maxIndex] {
				return -1
			}
		}
	}

	return maxIndex
}

// 这种方法非常精妙，我们只需要找到最大和第二大的值即可，如果最大的值比第二大的值至少大2倍，那么对于其他更小的值也是如此
// 因此一个循环就能搞定
func dominantIndex2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	largest, second := 0, 1

	if nums[largest] < nums[second] {
		largest, second = 1, 0
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > nums[largest] {
			largest, second = i, largest
		} else if nums[i] > nums[second] {
			second = i
		}
	}

	if nums[second]*2 <= nums[largest] {
		return largest
	}

	return -1
}
