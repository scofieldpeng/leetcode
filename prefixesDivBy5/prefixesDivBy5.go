package prefixesDivBy5

import (
	"strconv"
)

func prefixesDivBy5(A []int) []bool {
	length := len(A)
	res := make([]bool, length, length)
	count := 0
	for i, v := range A {
		count = count * 10 + v
		if k,_ := strconv.ParseInt(strconv.Itoa(count),2,64); k % 5== 0 {
			res[i] = true
		}
	}

	return res
}
