// problem: https://leetcode.com/problems/pairs-of-songs-with-total-durations-divisible-by-60/
// solution: https://algs.home.pjf.im/pairs-of-songs-with-total-durations-divisible-by-60.html
package numPairsDivisibleBy60

// 我们知道（x + y ) % 60 = 0,也等于 x % 60 + y % 60 = 60，也就是x % 60 = 60 - y % 60，但是因为y % 60有可能为0，但是x % 60不可能
// 为60，0-60怎么变为0-59呢？其实只需要再余一下60即可，也就是x % 60 = (60 - y % 60) % 60
// 而题目我们也知道，其实就是看有多少个(x+y) % 60=0，或者说，有多少对x % 60 = (60 - y % 60) % 60
func numPairsDivisibleBy60(time []int) int {
	res := 0
	mod := make([]int, 60, 60)
	for _, v := range time {
		// 这里其实非常精妙，其实就是看他这个(60-y%60)%60对应的x%60有多少个，其实如果按照暴力算法来看的话，也就是看这货前面有几个x%60=(60-y%60)%60
		// 这里的核心是，后面的和前面的进行比对，而不是前面的数字和后面的进行比对，这里是关键
		res += mod[(60-v%60)%60]
		mod[v%60]++
	}

	return res
}
