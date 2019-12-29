package isOneBitCharacter

func isOneBitCharacter(bits []int) bool {
	length := len(bits)
	if length < 1 {
		return false
	}
	if length == 1 && bits[0] == 0 {
		return true
	}

	isCharacter := true
	for i := 0; i <= length-2; i++ {
		if isCharacter {
			if bits[i] == 0 {
				continue
			}

			if bits[i] == 1 {
				isCharacter = false
				continue
			}
		} else {
			isCharacter = true
		}
	}

	if isCharacter {
		return true
	}

	return false
}

// 这种解法其实是解法1的优化版，原理也是一样的，从头遍历到倒数第二个
// 如果当前值是0，那么当前这个一定是一个单字符，跳下一位继续遍历
// 如果当前值是1，那么当前值一定是一个2bit的字符，那么跳2位继续遍历
func isOneBitCharacter2(bits []int) bool {
	i := 0
	length := len(bits)
	for ; i < length-1; {
		// 这里将i进行偏移，确定下次要遍历的数字
		i += bits[i] + 1
	}

	// 当遍历完后停留在最后一位，说明最后一位是一个单字符
	if i == length-1 {
		return true
	}

	return false
}

func isOneBitCharacter3(bits []int) bool {

}
