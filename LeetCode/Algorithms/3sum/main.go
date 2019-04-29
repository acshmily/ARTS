package main

import (
	"fmt"
	"sort"
)

/**
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	re := [][]int{}
	// 先排序
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			left, right, sum := i+1, len(nums)-1, 0-nums[i] // 定义双指针、以及检测值
			for left < right {
				if nums[left]+nums[right] == sum { //当匹配时
					re = append(re, []int{nums[left], nums[right], nums[i]}) //添加到结果数组里
					for (left < right) && (nums[left] == nums[left+1]) {     //去掉重复值
						left++
					}
					for (left < right) && (nums[right] == nums[right-1]) { // 去掉重复值
						right--
					}
					left++
					right--
				} else if nums[left]+nums[right] < sum { //左边的值小了
					for (left < right) && (nums[left] == nums[left+1]) {
						left++
					}
					left++
				} else { // 右边的值大了
					for (left < right) && (nums[right] == nums[right-1]) {
						right--
					}
					right--
				}
			}
		}

	}
	return re
}

func main() {
	fmt.Println(threeSum([]int{1, -1, -1, 0}))
}
