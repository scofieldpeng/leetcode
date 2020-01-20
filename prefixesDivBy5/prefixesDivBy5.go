package prefixesDivBy5

import "math"

func prefixesDivBy5(A []int) []bool {
	length := len(A)
	res := make([]bool, length, length)
	count := 0
	for i, v := range A {
		count += v * int(math.Mod(2, float64(i)))
		if count%5 == 0 {
			res[i] = true
		}
	}

	return res
}
