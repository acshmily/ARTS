package main

import (
	"fmt"
	"math"
)

/**
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

示例 1:

nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
示例 2:

nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m > n { // 将长度小的数组放在前面
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		tmp1 := m
		m = n
		n = tmp1
	}
	iMin, iMax, halfLen := 0, m, (m+n+1)/2 // 一半长度
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := halfLen - i
		if i < iMax && nums2[j-1] > nums1[i] {
			iMin = i + 1 // 太小
		} else if i > iMin && nums1[j-1] > nums2[i] {
			iMax = i - 1 // 太大
		} else {
			var maxLeft float64 = 0.0
			if i == 0 { // 当长度为2时
				maxLeft = float64(nums2[j-1])
			} else if j == 0 { // 当长度为1时
				maxLeft = float64(nums1[i-1])
			} else {
				maxLeft = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}
			if (m+n)%2 == 1 {
				return maxLeft
			}
			var minRight float64 = 0.0
			if i == m {
				minRight = float64(nums2[j])
			} else if j == n {
				minRight = float64(nums1[i])
			} else {
				minRight = math.Min(float64(nums2[j]), float64(nums1[j]))
			}

			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0

}
func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))

	fmt.Println(8 / 2)
}
