// problem: https://leetcode.com/problems/max-consecutive-ones/
// solution: https://algs.home.pjf.im/max-consecutive-ones/
package findMaxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	currentMax := 0
	for _, v := range nums {
		if v == 1 {
			currentMax++
		} else {
			currentMax = 0
			continue
		}
		if currentMax > res {
			res = currentMax
		}
	}

	return res
}
