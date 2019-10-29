// Given an array A of strings made only from lowercase letters, return a list of all characters that show up in all strings within the list (including duplicates).  For example, if a character occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.
//
// You may return the answer in any order.
//
//
//
// Example 1:
//
// Input: ["bella","label","roller"]
// Output: ["e","l","l"]
// Example 2:
//
// Input: ["cool","lock","cook"]
// Output: ["c","o"]
//
//
// Note:
//
// 1 <= A.length <= 100
// 1 <= A[i].length <= 100
// A[i][j] is a lowercase letter
package commanChars

import "math"

func commonChars(A []string) []string {
	// 因为题目要求了，只会是26个小写字母，因此定义一个26位长度的数组，然后将每个数组的值都设置为int16的最大值
	// （题目可以知道，每个字符的长度位100，所以，最大也不会超过100，所以uint8就够了，我看其他人直接用max int，其实没必要
	var totalChars [26]int
	for k := range totalChars {
		totalChars[k] = math.MaxUint8
	}

	// 返回值
	var res = make([]string, 0)

	// 开始遍历数组A
	for _, v := range A {
		// / 定义数组，用于判断该次循环的字符串中每个字符出现的次数
		var chars [26]int
		for _, vv := range []byte(v) {
			// / 这里需要记住的是，a的字符是97，所以，a-a=0,z-a=25
			chars[vv-'a']++
		}

		// / 遍历一下本次结果，然后其值和对应下标的值totalChars对比，然后取小的那个值，这里就取了某个字符出现的最小共同次数
		for i := 0; i < 26; i++ {
			if chars[i] < totalChars[i] {
				totalChars[i] = chars[i]
			}
		}
	}

	// / 遍历最后的结果，如果结果不为0并且不为初始的那个最大值，bingo，那么说明这个字符就是交集了，其value也是交集的次数
	for k, v := range totalChars {
		if v > 0 && v < math.MaxUint8 {
			for i := 0; i < v; i++ {
				// 这里的97+就是还原对应int的byte值，然后将其转化为string追加到结果数组中
				res = append(res, string(byte(97+k)))
			}
		}
	}

	return res
}
