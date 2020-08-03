// problem: https://leetcode.com/problems/binary-prefix-divisible-by-5
// solution: https://algs.home.pjf.im/95.html
package prefixesDivBy5

func prefixesDivBy5(A []int) []bool {
	length := len(A)
	res := make([]bool, length, length)
	count := 0
	for i, v := range A {
		// 这里运用了(ac+c)%d = ( (a%d)*(c%d) + c%d ) % d
		// 只关心余数，如果余数为0，那肯定整除
		count = (count * 2 + v) % 5
		if count % 5 == 0 {
			res[i] = true
		}
	}

	return res
}
