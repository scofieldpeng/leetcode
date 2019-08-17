package sortArrayByParityII

func sortArrayByParityII(A []int) []int {
	var result = make([]int, len(A))
	var oddIndex = 1
	var evenIndex = 0
	for _, v := range A {
		if v%2 == 0 {
			result[oddIndex] = v
			oddIndex++
		} else {
			result[evenIndex] = v
			evenIndex++
		}
	}

	return result
}
