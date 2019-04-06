package main

/**
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	//二分法
	low := 0
	high := len(nums)
	for low < high {
		mid := low + (high-low)/2 // 取中间值
		if nums[mid] > target {   // 中间值大了,则在分片的左边
			high = mid // 将边界改为 左边的最大值
		} else if nums[mid] == target { // 如果找到了,则返回索引位置
			return mid
		} else if nums[mid] < target { // 中间值小了,则分片在右边
			low = mid + 1 // 边界为中间值 + 1
		}
	}
	return low
}

func main() {
	searchInsert([]int{1, 3, 5, 6}, 2)
}
