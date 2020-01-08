// problem: https://leetcode.com/problems/x-of-a-kind-in-a-deck-of-cards/
// solution: https://algs.home.pjf.im/x-of-a-kind-in-a-deck-of-cards.html
package hasGroupsSizeX

// 核心是要确定x的值是多少
// 这里x的值其实就是min的值，如果所有的数值一样，那么说明正好能平分，如果不是，那看是否能转化为x * min
// 这里x其实是所有数字出现的最大公约数
func hasGroupsSizeX(deck []int) bool {
	groups := make(map[int]int)
	for _, v := range deck {
		if _, ok := groups[v]; !ok {
			groups[v] = 0
		}
		groups[v]++
	}

	x := 10000
	for _, v := range groups {
		if v < x {
			x = v
		}
	}

	if x < 2 {
		return false
	}

	x = 0
	// 如果有小于2的，直接返回false
	for _, v := range groups {
		x = gcd(x, v)
	}

	return x > 1
}

// 欧几里得算法（辗转相除法）求最大公约数
func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
