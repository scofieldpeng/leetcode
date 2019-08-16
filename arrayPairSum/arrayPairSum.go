package arrayPairSum

import (
	"sort"
)

func arrayPairSum(sum []int) int {
	if len(sum)%2 != 0 {
		panic("sum length should even")
	}

	sort.Ints(sum)
	numLen := len(sum) / 2
	res := 0
	for i := 0; i < numLen; i++ {
		if i > 0 {
			res = res + sum[i*2]
			continue
		}

		res += sum[0]
	}

	return res
}
