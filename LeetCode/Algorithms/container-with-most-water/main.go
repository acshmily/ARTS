package main

import "math"

/**
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.





The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].
In this case, the max area of water (blue section) the container can contain is 49.



Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

//面积为 底 * 高
//
func maxArea1(height []int) int { // 动态规划
	max, left, right := 0, 0, len(height)-1
	for left < right {
		// 比较每个面积
		max = int(math.Max(float64(max), math.Min(float64(height[left]), float64(height[right]))*float64(right-left)))
		if height[left] < height[right] { // 当左指针为短板则右移一位
			left++
		} else { // 又指针为短板则左移一位
			right--
		}
	}
	return max
}

func maxArea(height []int) int { // 时间复杂度存在效率问题
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ { // 计算每个面积
			max = int(math.Max(float64(max), math.Min(float64(height[i]), float64(height[j]))*float64(j-i)))
		}
	}
	return max
}

func main() {
	maxArea1([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
}
