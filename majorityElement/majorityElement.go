// problem: https://leetcode.com/problems/majority-element
// solution: https://algs.home.pjf.im/majority-element.html
package majorityElement

// hashmap solution
func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	numsMap := make(map[int]int)
	majorElement := nums[0]
	majorElementNum := 0
	for _, v := range nums {
		if _, exist := numsMap[v]; !exist {
			numsMap[v] = 0
		}
		numsMap[v]++

		if numsMap[v] > majorElementNum {
			majorElement = v
			majorElementNum = numsMap[v]
		}

		if majorElementNum > len(nums)/2 {
			return majorElement
		}
	}

	return 0
}
