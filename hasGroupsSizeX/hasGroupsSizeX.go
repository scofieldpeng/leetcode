package hasGroupsSizeX

import "sort"

func hasGroupsSizeX(deck []int) bool {
	sort.Ints(deck)

	groups := make(map[int]int)
	for _, v := range deck {
		if _, ok := groups[v]; !ok {
			groups[v] = 0
		}
		groups[v]++
	}

	return false
}
