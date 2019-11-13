package twoSumII

func twoSum(numbers []int, target int) []int {
	leftIndex := 0
	rightIndex := len(numbers) - 1

	for ; leftIndex < rightIndex; {
		sum := numbers[leftIndex] + numbers[rightIndex]
		if sum == target {
			return []int{leftIndex + 1, rightIndex + 1}
		}
		if sum > target {
			rightIndex--
		} else {
			leftIndex++
		}

	}

	return []int{}
}
