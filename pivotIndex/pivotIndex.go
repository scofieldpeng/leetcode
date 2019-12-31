// problem: https://leetcode.com/problems/find-pivot-index/
// solution: https://algs.home.pjf.im/find-pivot-index.html
package pivotIndex

// 假设数组总数为total，然后pivot的值是x，然后该pivot到最左侧的值是a，privot到最右侧的值是b，那么可以知道
// a + b - y = total
// 又因为a = b,所以
// 2a - y = total
// 所以我们只要找到符合这个等式的值即可
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	a := 0
	for i := 0; i < len(nums); i++ {
		a += nums[i]
		if a*2 - nums[i] == total {
			return i
		}
	}

	return -1
}
