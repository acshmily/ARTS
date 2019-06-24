package main

import "sort"

/**
给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。
*/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
