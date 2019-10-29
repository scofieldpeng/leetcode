// Let's define a function f(s) over a non-empty string s, which calculates the frequency of the smallest character in s. For example, if s = "dcce" then f(s) = 2 because the smallest character is "c" and its frequency is 2.
//
// Now, given string arrays queries and words, return an integer array answer, where each answer[i] is the number of words such that f(queries[i]) < f(W), where W is a word in words.
//
//
//
// Example 1:
//
// Input: queries = ["cbd"], words = ["zaaaz"]
// Output: [1]
// Explanation: On the first query we have f("cbd") = 1, f("zaaaz") = 3 so f("cbd") < f("zaaaz").
// Example 2:
//
// Input: queries = ["bbb","cc"], words = ["a","aa","aaa","aaaa"]
// Output: [1,2]
// Explanation: On the first query only f("bbb") < f("aaaa"). On the second query both f("aaa") and f("aaaa") are both > f("cc").
//
//
// Constraints:
//
// 1 <= queries.length <= 2000
// 1 <= words.length <= 2000
// 1 <= queries[i].length, words[i].length <= 10
// queries[i][j], words[i][j] are English lowercase letters.
package numSmallerByFrequency

import (
	"fmt"
	"sort"
)

// 这道题首先计算出每个字符串中最小字母出现的次数，然后，queries的每一个下标的结果集再和words的结果集相比较，得出qeries比words的结果集
// 的数量低的一个int数组集
func numSmallerByFrequency(queries []string, words []string) []int {
	queriesCount := make([]int, 0, len(queries))
	wordsCount := make([]int, 0, len(words))
	for _, v := range queries {
		queriesCount = append(queriesCount, smallestCharacterNum(v))
	}
	for _, v := range words {
		wordsCount = append(wordsCount, smallestCharacterNum(v))
	}

	sort.Ints(wordsCount)

	fmt.Println(queriesCount)
	fmt.Println(wordsCount)

	resCount := make([]int, len(queries), len(queries))
	for k, v := range queriesCount {
		for kk, vv := range wordsCount {
			if vv > v {
				resCount[k] = len(wordsCount) - kk
				break
			}
		}
	}

	return resCount
}

// / 查找最小的字符串出现的次数
func smallestCharacterNum(character string) int {
	var res = 1
	var smallByte byte = 'z'
	var length = len(character)
	for i := 0; i < length; i++ {
		if character[i] <= smallByte {
			if character[i] < smallByte {
				res = 1
				smallByte = character[i]
			}
			res++
		}
	}

	return res
}

// 这个是我看到非常精妙的答案,精妙的地方在于我上面的解法的时间复杂度是2n+n2,并且这里还是没有包括对words的结果集进行排序的时间复杂度
// 而这里，他与我的恰恰相反，先求出words的counter结果集，然后此时通过一个循环，能够计算出对于任意一个范围内的数组，比他大的数字有多少个
func numSmallerByFrequency2(queries []string, words []string) []int {

	lookup := make([]int, 11)
	for _, w := range words {
		n := counter(w)
		// 精妙的在于这里，对于每次求出的结果都放到lookup这个变量，这里最开始不太明白，明白了之后不得不鼓下掌，
		// 举个例子就明白了，如果是对于1，2，3，4，5来说，如果是1，那么比1大于的值是4，对于2来说，值为3，因此，如果是连续的值，对于当前值
		// 来说就是，最大的那个值减去当前值，那么，如果不是连续的呢？我们只需要知道，对于比他大的数字，出现过多少次就好了
		// 比如我们这里只有1，2，3，5，那么对于5来说，对于1来说的话是大于的，对于2，3也是如此，因此就在这几个数字的头上增加1，这样后面我
		// 们就能直接查看一个对应数字的大于的数量了，非常精妙，当然，这里的限制条件就是，只有最小值的最长的数量是固定的才可以，否则反而还会
		// 增加查询次数
		for j := n; j > 0; j-- {
			// 这里为什么要j-1呢，是因为数组下标是从0开始的
			lookup[j-1]++
		}
	}

	result := make([]int, len(queries))
	for i, q := range queries {
		n := counter(q)
		result[i] = lookup[n]
	}
	fmt.Println(lookup)
	return result
}

// 他这里和我的解法也不同，他是生成了一个26位长度的数组，然后将对应字母出现的次数存储在对应下标，然后只需要获取到最小的字母即可
// 当然，用数组的解法可以这样，但是我认为没必要，浪费了26个字节的数组空间
func counter(s string) int {
	si := 27
	index := [26]int{}

	for i := 0; i < len(s); i++ {
		c := int(s[i] - 97)
		if si > c {
			si = c
		}
		index[c]++
	}

	return index[si]
}
