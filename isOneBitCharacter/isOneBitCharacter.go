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

// 贪婪算法，虽然我看懂了solution的算法，但实际上我没有搞懂这跟贪婪算法有什么关系
//
// 我们仔细思考规则可以发现，如果一个字符是1，那么下一位不管是1还是0，下一位都是一个双位字符，也就是说，11，111x都是一个双位
// 那么，我们就从右往左找最近的那个0（不包含最后一位数的那个0），只要找到了，然后看在这个0和最后一位之间间隔的1是多少位，如果是一个奇数位的
// 话，那么就说明，最后一位的0肯定是这个1的组成部分，如果是偶数位，那么说明，这之间的1肯定都已经互相组合了
func isOneBitCharacter3(bits []int) bool {
	findOneCount := 0
	for start := len(bits) - 2; start >= 0; start-- {
		// 找1的数量，这里还可以用位运算, 但是对于go来说有点多此一举了
		// bits[start] ^ 1 == 0
		if bits[start] == 1 {
			findOneCount++
			continue
		}
		// 一旦找到0就break
		break
	}

	// 如果找到的1是偶数，那么最后一位肯定是一个独位0
	return findOneCount%2 == 0
}
