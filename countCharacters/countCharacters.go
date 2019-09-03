// You are given an array of strings words and a string chars.
//
// A string is good if it can be formed by characters from chars (each character can only be used once).
//
// Return the sum of lengths of all good strings in words.
//
//
//
// Example 1:
//
// Input: words = ["cat","bt","hat","tree"], chars = "atach"
// Output: 6
// Explanation:
// The strings that can be formed are "cat" and "hat" so the answer is 3 + 3 = 6.
// Example 2:
//
// Input: words = ["hello","world","leetcode"], chars = "welldonehoneyr"
// Output: 10
// Explanation:
// The strings that can be formed are "hello" and "world" so the answer is 5 + 5 = 10.
//
//
// Note:
//
// 1 <= words.length <= 1000
// 1 <= words[i].length, chars.length <= 100
// All strings contain lowercase English letters only.
package countCharacters

import "fmt"

func countCharacters(words []string, chars string) int {
	var res = 0

	// 题目规定了只可能有26个小写字母，因此可以用数组的0-25来代表每个字母，其值为该字母出现的次数
	var charsArr [26]int
	for _, v := range []byte(chars) {
		charsArr[v-'a']++
	}

	for _, v := range words {
		if isContainCharacters(v, charsArr) {
			fmt.Println("v:" + v + " is safe")
			res = len(v) + res
		}
	}

	return res
}

// 查找一个单词是的所有字符否在对应字符里面，如果存在，返回true
// 这里解决思路其实就是遍历words，因为chars存储了我们字符集每个出现的次数，所以如果遇到就将chars对应的值减1，如果<1,那么说明该字母已经
// 不在我们的字母里面了，此时直接退出，索命这个单词不是安全字母
func isContainCharacters(word string, chars [26]int) bool {
	var isOk = true

	for _, v := range []byte(word) {
		if chars[v-'a'] > 0 {
			chars[v-'a']--
			continue
		}
		isOk = false
		break
	}

	return isOk
}
