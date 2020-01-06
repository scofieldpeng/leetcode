package maxDistToClosest

// 看起来又是一道dp问题
// 找到最长的一段空位，然后找到中间的位置或者末尾/开头的位置就是最大的了
func maxDistToClosest(seats []int) int {
	length := len(seats)

	// 1. 找到最长空位的开始和结束位置
	// 2. 如果最长的那截开始是0，那么坐在0，如果结束是length-1,那么坐在length-1
	//	  其他情况选中间的位置
	max := 0
	current := 0
	start := 0

	for i := 0; i < length; i++ {
		if seats[i] == 0 {
			if current == 0 {
				start = i
			}
			current = i - start + 1
			if start != 0 && i != length-1 {
				current = current/2 + current%2
			}
		}

		if current > max {
			max = current
		}

		if seats[i] == 1 {
			current = 0
		}

	}

	if max == 0 {
		return 1
	}

	return max
}
