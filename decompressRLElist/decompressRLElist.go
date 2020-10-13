package decompressRLElist

func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		freq, value := nums[i], nums[i+1]
		res = append(res, generateValue(value, freq)...)
	}

	return res
}

func generateValue(num, loop int) []int {
	res := make([]int, loop, loop)
	for i := 0; i < loop; i++ {
		res[i] = num
	}

	return res
}
