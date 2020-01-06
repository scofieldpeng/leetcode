// problem: https://leetcode.com/problems/positions-of-large-groups/
// solution: https://algs.home.pjf.im/positions-of-large-groups.html
package largeGroupPositions

func largeGroupPositions(S string) [][]int {
	length := len(S)
	start := 0
	res := make([][]int, 0)

	for i := 0; i < length; i++ {
		if i == length-1 || S[i] != S[i+1] {
			if i-start+1 >= 3 {
				res = append(res, []int{start, i})
			}
			start = i + 1
		}
	}

	return res
}
