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
	length := len(nums)
	res := make([]int, 0, length)

	for i := 0; i < length; i++ {
		//  + length是为了确保这个下标对应的值已经存在了
		//  % length是为了能够哪个下标对应的那个原来的值
		nums[ (nums[i]-1)%length ] += length
	}

	// 这里遍历是为了找到上面遍历过程中没有对应值的哪个下标，这样就知道对应下标的那个值是丢失的了
	for i := 0; i < length; i++ {
		if nums[i] <= length {
			res = append(res, i+1)
		}
	}

	return res
}
