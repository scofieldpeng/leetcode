// problem: https://leetcode.com/problems/plus-one
// solution: https://algs.home.pjf.im/plus-one
package plusOne

func plusOne(digits []int) []int {
	var (
		shouldStepForward = false
		length            = len(digits)
	)
	digits[length-1] += 1
	for i := length - 1; i >= 0; i-- {
		if shouldStepForward {
			digits[i] += 1
			shouldStepForward = false
		}
		if digits[i] > 9 {
			shouldStepForward = true
			digits[i] = digits[i] - 10
		}
	}

	if shouldStepForward {
		return append([]int{1}, digits...)
	}

	return digits
}
