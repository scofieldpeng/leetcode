// problem: https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
package findSpecialInteger

func findSpecialInteger(arr []int) int {
	var count int = 0

	for i := 0; i < len(arr) - 1; i++ {
		if arr[i+1] == arr[i] {
			count++
		} else {
			count = 1
		}

		if len(arr)/count < 4 {
			return arr[i]
		}
	}

	return arr[0]
}
