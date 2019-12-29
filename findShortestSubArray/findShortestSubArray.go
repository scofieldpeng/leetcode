// problem: https://leetcode.com/problems/degree-of-an-array/
// solution: https://algs.home.pjf.im/degree-of-an-array.html
package findShortestSubArray

type dataFrequency struct {
	Start int
	End   int
}

func findShortestSubArray(nums []int) int {
	numLength := make(map[int]int)
	numFreInfo := make(map[int]dataFrequency)
	length := len(nums)
	degree := 0
	// 找到每一个数字出现的频率，第一次出现的下标和最后一次的下标
	// 同时，我们还能找到degree是多少
	for i := 0; i < length; i++ {
		if _, ok := numLength[nums[i]]; !ok {
			numLength[nums[i]] = 0
			numFreInfo[nums[i]] = dataFrequency{
				Start: i,
				End:   0,
			}
		}
		numLength[nums[i]]++
		numFreInfo[nums[i]] = dataFrequency{
			Start: numFreInfo[nums[i]].Start,
			End:   i,
		}
		if degree <= numLength[nums[i]] {
			degree = numLength[nums[i]]
		}
	}

	// 知道degree后，找到最短的那个
	res := length
	for k, v := range numLength {
		if v == degree {
			res = min(res, numFreInfo[k].End-numFreInfo[k].Start+1)
		}
	}

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
