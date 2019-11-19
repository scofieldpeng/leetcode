// problem: https://leetcode.com/problems/contains-duplicate-ii
// solution: https://algs.home.pjf.im/contains-duplicate-ii.html
package containsDuplicateII

func containsNearbyDuplicate(nums []int, k int) bool {
	numsMap := make(map[int]int, len(nums))
	for index, v := range nums {
		if vv, exist := numsMap[v]; exist {
			if index - vv <= k {
				return true
			}
		}
		numsMap[v] = index
	}

	return false
}
