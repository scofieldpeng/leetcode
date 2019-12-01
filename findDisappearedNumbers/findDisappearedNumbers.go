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

func findDisappearedNumbers3(nums []int) []int {
	length := len(nums)
	res := make([]int, 0, length)
	for i := 0; i < length; i++ {
		// 找到对应索引值对应的真正的索引值
		index := abs(nums[i])
		if nums[index-1] > 0 {
			// 要不变为0（恰好在本身的位置）
			// 要不变为负数(不在本身位置，此时负数的绝对值是该位置上的那个值本身)
			nums[index-1] = -nums[index-1]
		}
	}

	// 如果值大于0，所以该位置没有人占用，那么说明这一位是disappear的
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}

	return res
}

func abs(data int) int {
	if data < 0 {
		return -data
	}
	return data
}
