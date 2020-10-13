// problem: https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/
package sumZero

// a + (-a) = 0
// 如果是奇数，那么n中落单的为0即可满足
func sumZero(n int) []int {
	res := make([]int, n, n)
	i := 0
	if n%2 == 1 {
		i = 1
	}
	start := 1
	for ; i < n; i += 2 {
		res[i] = start
		res[i+1] = -start
		start++
	}

	return res
}
