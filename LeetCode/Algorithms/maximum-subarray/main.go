package main

import (
	"fmt"
)

/**
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

*/

func maxSubArray(nums []int) int {
	sum, maxSum := 0, nums[0]
	for i := 0; i < len(nums); i++ {
		if sum > 0 { // 如果当前结果大于0则继续往后相加
			sum = sum + nums[i]
		} else { // 如果当前结果小于等于0 ,则当前结果等于当前元素值
			sum = nums[i]
		}
		if sum > maxSum { // 如果当前结果大于最大值,则赋值
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}
