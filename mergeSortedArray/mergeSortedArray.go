// problem: https://leetcode.com/problems/merge-sorted-array
// solution: https://algs.home.pjf.im/merge-sorted-array
package mergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	resLastIndex := m + n - 1
	for ; m > 0 && n > 0; resLastIndex-- {
		if nums1[m-1] >= nums2[n-1] {
			nums1[resLastIndex] = nums1[m-1]
			m -= 1
			continue
		}

		nums1[resLastIndex] = nums2[n-1]
		n -= 1
	}

	for ; n > 0; n-- {
		nums1[resLastIndex] = nums2[n-1]
		resLastIndex--
	}
}
