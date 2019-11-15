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
	fmt.Printf("遍历元素开始,开始:%d,结束:%d\n", low, high)
	if low == high {
		fmt.Printf("遍历一个元素，直接返回，遍历的元素是nums[%d]=%d\n", low, nums[low])
		return nums[low]
	}

	mid := (high-low)/2 + low
	fmt.Printf("计算本轮的开始%d,中间:%d,结束:%d\n", low, mid, high)
	left := majorElementRec(nums, low, mid)
	fmt.Printf("本轮计算后开始(%d)和中间(%d)的胜出结果是%d\n", low, mid, left)
	right := majorElementRec(nums, mid+1, high)
	fmt.Printf("本轮计算后中间+1(%d)和结尾(%d)的胜出结果是%d\n", mid+1, high, right)

	if left == right {
		fmt.Printf("左右两边(%d:%d)的结果一致,返回结果(开始:%d,结束:%d)\n", left, right, low, high)
		return left
	}

	leftCount := elementCount(nums, left, low, high)
	rightCount := elementCount(nums, right, low, high)

	fmt.Printf("计算左边的最大值(%d)和右边的最大值(%d)在nums[%d]-nums[%d]中的出现次数%d:%d\n", left, right, low, high, leftCount, rightCount)

	if leftCount >= rightCount {
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

// BoyerMoore vote solution
// see: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
// or can see: http://www.cs.utexas.edu/~moore/best-ideas/mjrty/example.html
func majorElementByBoyerMooreVoteAlgs(nums []int) int {
	num := 0
	count := 0
	for _, v := range nums {
		if count == 0 {
			num = v
		}
		if num == v {
			count++
		} else {
			count--
		}
	}

	return num
}
