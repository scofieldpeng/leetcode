package imageSmoother

func imageSmoother(M [][]int) [][]int {
	res := make([][]int, 0, len(M))
	yLen := len(M[0])
	xLen := len(M)
	count := 0
	// x为纵坐标，y为横坐标
	for x, v := range M {
		res = append(res, make([]int, yLen, yLen))
		for y := range v {
			count = 0
			// 遍历其周围8个方向的数据，如果存在，那么就将其
			for k := -1; k <= 1; k++ {
				for j := -1; j <= 1; j++ {
					// 当发现他周围的数量的时候，增加其数量
					if k+x >= 0 && k+x < xLen && j+y >= 0 && j+y < yLen {
						count++
						res[x][y] += M[x+k][y+j]
					}
				}
			}
			if count > 0 {
				res[x][y] = res[x][y] / count
			}
		}
	}

	return res
}
