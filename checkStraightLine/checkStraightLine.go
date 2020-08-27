// problem: https://leetcode.com/problems/check-if-it-is-a-straight-line
package checkStraightLine

// 两条线就可以确定其在坐标系的正切值(tan),然后不断判断即可
// 但是有两个例外要处理下，那就是水平和垂直的情况
func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 2 {
		return false
	}
	if len(coordinates[0]) != 2 {
		return false
	}
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		if len(coordinates[i]) != 2 {
			return false
		}
		x, y := coordinates[i][0], coordinates[i][1]

		// 水平和垂直的情况
		if (x1 == x2 && x2 == x) || (y1 == y2 && y2 == y) {
			continue
		}

		if float64(x1-x)/float64(y1-y) != float64(x2-x)/float64(y2-y) {
			return false
		}
	}

	return true
}
