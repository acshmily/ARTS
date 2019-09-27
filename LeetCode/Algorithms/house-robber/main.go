package main

import (
	"fmt"
	"math"
)

/**
动态规划
*/
func rob(nums []int) int {
	pre := 0
	current := 0
	for _, v := range nums {
		temp := current                                           // 临时变量存储当前值
		current = int(math.Max(float64(pre+v), float64(current))) // 获取当前最大值
		pre = temp                                                // 移动标记
	}
	return current
}

func main() {
	fmt.Print(rob([]int{2, 7, 9, 3, 1}))
}
