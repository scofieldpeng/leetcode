package minCostToMoveChips

// 移动偶数步数不需要消耗点数，移动单数步数，每移动一个需要一个点数
// 也就是，可以把所有的筹码都放置到最中间，然后看最中间的两，结果就是最低的那两位了
func minCostToMoveChips(position []int) int {
	oddChips := 0
	evenChips := 0

	if len(position) <= 1 {
		return 0
	}
	if len(position) == 2 {
		return 1
	}

	for _, v := range position {
		if v%2 == 0 {
			evenChips++
		} else {
			oddChips++
		}
	}

	if evenChips >= oddChips {
		return oddChips
	}

	return evenChips
}
