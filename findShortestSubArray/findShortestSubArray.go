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
	res := length
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

		// 这里我们不断找数组的degree，由于我们每次在确定degree的时候，意味着，我们最后要找的答案也基本在这里
		// 如果我们发现了更大的degree，说明，当前的数组如果是最后一次，这个就是要找的最小数组了
		// 如果我们发现数组的degree和我们目前已经找到的degree相同，那么我们就找这两个有差异的degree的数组的最小数组即可
		if numLength[nums[i]] > degree {
			res = numFreInfo[nums[i]].End - numFreInfo[nums[i]].Start + 1
			degree = numLength[nums[i]]
			continue
		}

		if numLength[nums[i]] == degree {
			res = min(res, numFreInfo[nums[i]].End-numFreInfo[nums[i]].Start+1)
			degree = numLength[nums[i]]
			continue
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
