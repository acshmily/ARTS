package main

import "fmt"

/**
  输入: numbers = [2, 7, 11, 15], target = 9
  输出: [1,2]
  解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

*/
// 暴力破解法
func twoSum1(numbers []int, target int) []int {
	re := make([]int, 0)
	for i := 0; i < len(numbers); i++ {
		// 出口
		if numbers[i] > target {
			return nil
		}
		// 获取剩余值
		tmp := target - numbers[i]
		for x := i + 1; x < len(numbers); x++ {
			if numbers[x] == tmp {
				return []int{i + 1, x + 1}
			} else if numbers[x] > tmp {
				break
			}
		}
	}
	return re
}

// 双指针
func twoSum(numbers []int, target int) []int {
	listEndPoint := len(numbers) - 1
	re := make([]int, 0)
	if len(numbers) <= 1 {
		return re
	}
	i := 0
	for i >= 0 && i < len(numbers) && listEndPoint > 0 && listEndPoint < len(numbers) {
		// 定义出口
		if numbers[i]+numbers[listEndPoint] == target {
			return []int{i + 1, listEndPoint + 1}
			// 大了,尾部指针往前移动
		} else if numbers[i]+numbers[listEndPoint] > target {
			listEndPoint--
			// 小了,头部指针往后移动
		} else {
			i++
		}

	}
	return re

}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
