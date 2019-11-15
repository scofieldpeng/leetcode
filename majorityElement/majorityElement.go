// problem: https://leetcode.com/problems/majority-element
// solution: https://algs.home.pjf.im/majority-element.html
package majorityElement

import "fmt"

// hashmap solution
func majorityElement(nums []int) int {
	if len(nums) < 1 {
		return 0
	}

	numsMap := make(map[int]int)
	majorElement := nums[0]
	majorElementNum := 0
	for _, v := range nums {
		if _, exist := numsMap[v]; !exist {
			numsMap[v] = 0
		}
		numsMap[v]++

		if numsMap[v] > majorElementNum {
			majorElement = v
			majorElementNum = numsMap[v]
		}

		if majorElementNum > len(nums)/2 {
			return majorElement
		}
	}

	return 0
}

func majorElementByDivideAndConquer(nums []int) int {
	return majorElementRec(nums, 0, len(nums)-1)
}

func majorElementRec(nums []int, low, high int) int {
	if low == high {
		fmt.Printf("marjorElementRec low equal high,return, low: %d, high: %d \n", low, high)
		return nums[low]
	}

	mid := (high-low)/2 + low
	fmt.Printf("marjorElementrec low: %d, hight: %d, mid: %d \n", low, high, mid)
	left := majorElementRec(nums, low, mid)
	right := majorElementRec(nums, mid+1, high)

	if left == right {
		fmt.Printf("majorElementRec left==right, left: %d, right: %d \n", left, right)
		return nums[left]
	}

	leftCount := elementCount(nums, left, low, high)
	fmt.Printf("majorElementRec call elementCount left:%d, low:%d, high:%d,leftCount:%d\n", left, low, high, leftCount)
	rightCount := elementCount(nums, right, low, high)
	fmt.Printf("majorElementRec call elementCount right:%d, low:%d, high:%d,leftCount:%d\n", right, low, high, rightCount)

	if leftCount > rightCount {
		return left
	}

	return right
}

// 获取某个数字在某区间的出现次数
func elementCount(nums []int, number int, start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if nums[i] == number {
			count++
		}
	}

	return count
}
