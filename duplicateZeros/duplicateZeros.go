// Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.
//
// Note that elements beyond the length of the original array are not written.
//
// Do the above modifications to the input array in place, do not return anything from your function.
//
//
//
// Example 1:
//
// Input: [1,0,2,3,0,4,5,0]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]
// Example 2:
//
// Input: [1,2,3]
// Output: null
// Explanation: After calling your function, the input array is modified to: [1,2,3]
//
//
// Note:
//
// 1 <= arr.length <= 10000
// 0 <= arr[i] <= 9
package duplicateZeros

import "fmt"

func duplicateZeros(arr []int) {
	zeroNumbers, firstZeroIsDuplicate := duplicateZeroNumbers(arr)
	startLastIndex := len(arr) - 1
	// 只有大于0的才需要位移
	if zeroNumbers > 0 {
		// 确定开始位移的最后一个index
		index := startLastIndex
		isStart := true
		for startLastIndex = startLastIndex - zeroNumbers; startLastIndex >= 0 && index >= 0 && zeroNumbers > 0; startLastIndex-- {
			arr[index] = arr[startLastIndex]
			index = index - 1
			if arr[startLastIndex] == 0 && index > 0 {
				if !firstZeroIsDuplicate && isStart {
					isStart = false
					continue
				}
				arr[index] = arr[startLastIndex]
				index = index - 1
				// 这里要记住，最多位移这么多位，移动完之后，剩下的就算有0也不移动了
				zeroNumbers--
			}
			isStart = false
		}
	}
}

// 查找数组里面应该有多少个重复的0，注意，如果数组最后一个是0，那么跳过，第二个参数返回的是，到时候要跳过的最后一个0是否duplicate，这里
// 主要是防止边界条件的0
func duplicateZeroNumbers(arr []int) (int, bool) {
	var searchLastIndex = len(arr) - 1
	var res = 0
	var zeroNumbers = 0
	for i := 0; i <= searchLastIndex-res; i++ {
		if arr[i] == 0 {
			zeroNumbers++
		}
		if arr[i] == 0 && i != searchLastIndex-res {
			res++
		}
	}

	fmt.Printf("source %v, numbers: %d, zeroNumbers:%d\n", arr, res, zeroNumbers)

	return res, zeroNumbers == res
}
