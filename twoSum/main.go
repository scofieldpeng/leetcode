/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9, Because nums[0] + nums[1] = 2 + 7 = 9, return [0, 1].
*/
package main

func twoSum(nums []int, target int) []int {
	length := len(nums)
	res := make([]int, 2)
	// find
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if target == nums[i]+nums[j] {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	// not found
	res = make([]int, 0)
	return res
}
