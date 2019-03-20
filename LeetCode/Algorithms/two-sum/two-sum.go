package main

import "fmt"

/**
Q:
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func twoSum(nums []int, target int) []int {
	// 生成一个Map<Key为 数组的值,Value为索引值>
	valMap := make(map[int]int, len(nums))
	// 循环输入存入Map , 降低时间复杂度
	for key, value := range nums {
		re := target - value   //目标值 - 数组索引第一个 值
		int2, ok := valMap[re] // 获取索引值
		if ok {                //如果存在
			if int2 != key {
				return []int{key, int2}
			}
		}
		valMap[value] = key
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 18
	re := twoSum(nums, target)
	fmt.Printf("%v", re)
}
