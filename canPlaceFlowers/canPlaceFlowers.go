package canPlaceFlowers

// 只需要保证它本身是0，且左右两边也为0那么就意味着它可以种花
// 需要注意的是，首尾的话只需要判断一侧即可
func canPlaceFlowers(flowerbed []int, n int) bool {
	var canPlantNum = 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			// 种上一朵小花
			flowerbed[i] = 1
			// 可以栽的花多一棵
			canPlantNum++
		}
	}

	if canPlantNum >= n {
		return true
	}

	return false
}
