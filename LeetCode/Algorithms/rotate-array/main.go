package main

import "fmt"

/**
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
*/
func rotate(nums []int, k int) {
	if k == 0 || len(nums) == 0 {
		return
	}
	k = k % len(nums)
	if k == 0 {
		return
	}
	tmp := nums
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	for i := 0; i < len(tmp); i++ {
		tmp[i] = nums[i]
	}
	return

}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(list, 3)
	fmt.Println(list)
}
