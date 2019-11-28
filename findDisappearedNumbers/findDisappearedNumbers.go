// problem: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/
// solution: https://algs.home.pjf.im/find-all-numbers-disappeared-in-an-array.html
package findDisappearedNumbers

func findDisappearedNumbers(nums []int) []int {
	tmp := make([]int, len(nums), len(nums))
	for _, v := range nums {
		tmp[v-1] = v
	}
	res := make([]int, 0)
	for k, v := range tmp {
		if v == 0 {
			res = append(res, k+1)
		}
	}

	return res
}

func findDisappearedNumbers2(nums []int) []int {
	return []int{}
}
