package heightChecker

import "sort"

func heightChecker(heights []int) int {
	if len(heights) > 100 {
		panic("height length should not exceed than 100")
	}
	source := make([]int, len(heights))
	copy(source, heights)
	sort.Ints(heights)

	res := 0
	for k, v := range heights {
		if source[k] != v {
			res++
		}
	}

	return res
}
