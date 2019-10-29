// For the i-th query val = queries[i][0], index = queries[i][1], we add val to A[index].  Then, the answer to the i-th query is the sum of the even values of A.
//
// (Here, the given index = queries[i][1] is a 0-based index, and each query permanently modifies the array A.)
//
// Return the answer to all queries.  Your answer array should have answer[i] as the answer to the i-th query.
//
//
//
// Example 1:
//
// Input: A = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
// Output: [8,6,2,4]
// Explanation:
// At the beginning, the array is [1,2,3,4].
// After adding 1 to A[0], the array is [2,2,3,4], and the sum of even values is 2 + 2 + 4 = 8.
// After adding -3 to A[1], the array is [2,-1,3,4], and the sum of even values is 2 + 4 = 6.
// After adding -4 to A[0], the array is [-2,-1,3,4], and the sum of even values is -2 + 4 = 2.
// After adding 2 to A[3], the array is [-2,-1,3,6], and the sum of even values is -2 + 6 = 4.
//
//
// Note:
//
// 1 <= A.length <= 10000
// -10000 <= A[i] <= 10000
// 1 <= queries.length <= 10000
// -10000 <= queries[i][0] <= 10000
// 0 <= queries[i][1] < A.length
package sumEvenAfterQueries

// 这道题也比较简单，其实就是queries的子数组的1下标为A的下标，0下标为对应A下标要增加的值，然后操作完成后，计算一下A中最新的偶数的数量
// 简单粗暴更新A数组后，重新计算一下获取数组A的偶数和
// 但是，这样每次都要计算数组A的偶数和，效率太低了，为啥呢，因为每次只会更新数组A的一个值，完全没必要每次重新计算A的偶数和，只需要最开始计算
// 一次，然后盘更新的这个是是否为偶数即可，情况共有4种：
// 1. 要更新的数是偶数
//   1) 更新完后是偶数，此时计算两次偶数的数值差，然后和上一次的偶数和相加就是新的偶数和
//   2) 更新完后是奇数，将上次的偶数和的值减去本次更新的那个偶数就是新的偶数和
// 2. 要更新的数是奇数
//   3) 更新完后是奇数，新的偶数和和上次一致，没有更新
//   4) 更新完后是偶数，新的偶数和就是上次的偶数和加上本次更新后的偶数
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	var res = make([]int, len(queries), len(queries))
	var prevEvenNum = 0
	var oldValue = 0

	for _, v := range A {
		if v&1 == 0 {
			prevEvenNum += v
		}
	}

	for k, v := range queries {
		oldValue = A[v[1]]
		A[v[1]] += v[0]
		if oldValue&1 == 0 {
			// case 1
			if A[v[1]]&1 == 0 {
				prevEvenNum += v[0]
				// case 2
			} else {
				prevEvenNum -= oldValue
			}
		} else {
			// case 4
			if A[v[1]]&1 == 0 {
				prevEvenNum += A[v[1]]
			}
		}

		res[k] = prevEvenNum
	}

	return res
}
