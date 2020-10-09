// problem: https://leetcode.com/problems/find-numbers-with-even-number-of-digits
package findNumbers

import "math"

func findNumbers(nums []int) int {
	res := 0
	for _, v := range nums {
		tmp := getEvenNumber(v)
		if tmp&1 == 0 {
			res++
		}
	}

	return res
}

func getEvenNumber(data int) int {
	return int(math.Log10(float64(data))) + 1
}
