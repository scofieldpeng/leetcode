package numEquivDominoPairs

func numEquivDominoPairs(dominoes [][]int) int {
	length := len(dominoes)
	res := 0
	for i := 0; i < length-1; i++ {
		for j := i+1; j < length; j++ {
			if (dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1]) ||
				(dominoes[i][1] == dominoes[j][0] && dominoes[i][0] == dominoes[j][1]) {
				res++
			}
		}
	}

	return res
}
