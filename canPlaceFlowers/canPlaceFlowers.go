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

// 第一种算法的改良版
// 1. 找到够n的就停止，不用继续了
// 2. 如果i上能种树，那么i+1是不能种的，跳过
func canPlaceFlowers2(flowerbed []int, n int) bool {
	var canPlantNum = 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			// 种上一朵小花
			flowerbed[i] = 1
			// 可以栽的花多一棵
			canPlantNum++
			// 这里可以加速一下，因为下一位肯定不能栽了
			i++
		}
		if canPlantNum >= n {
			return true
		}
	}

	return false
}
