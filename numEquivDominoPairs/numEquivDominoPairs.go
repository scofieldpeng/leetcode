package numEquivDominoPairs

func numEquivDominoPairs(dominoes [][]int) int {
	length := len(dominoes)
	res := 0
	for i := 0; i < length-1; i++ {
		if (dominoes[i][0] == dominoes[i+1][0] && dominoes[i][1] == dominoes[i+1][1]) ||
			(dominoes[i][1] == dominoes[i+1][0] && dominoes[i][0] == dominoes[i+1][1]) {
			res++
		}
	}

	return res
}
