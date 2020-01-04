// problem: https://leetcode.com/problems/positions-of-large-groups/
// solution: https://algs.home.pjf.im/positions-of-large-groups.html
package largeGroupPositions

func largeGroupPositions(S string) [][]int {
	length := len(S)
	start := 0
	prev := S[0]
	res := make([][]int, 0)

	for i := 1; i < length-1; i++ {
		if S[i] != prev {
			// i -1 - start + 1 = i - start
			if i-start >= 3 {
				res = append(res, []int{start, i - 1})
			}
			start = i
			prev = S[i]
			continue
		}
	}

	return res
}
