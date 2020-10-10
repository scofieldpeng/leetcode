// problem: https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
package replaceElements

func replaceElements(arr []int) []int {
	max := -1
	res := make([]int, len(arr), len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		// 如果发现当前值比max还大，那么交换
		if arr[i] > max {
			max, res[i] = arr[i], max
			continue
		}
		// 替换掉
		if arr[i] <= max {
			res[i] = max
		}
	}

	return res
}
