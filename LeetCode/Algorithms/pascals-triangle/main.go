package main

/**
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/
func generate(numRows int) [][]int {
	re := [][]int{}
	var gofxck func(num int)
	gofxck = func(num int) {
		if num == numRows { // 递归出口
			return
		}
		if num >= len(re) { // 初始化容量
			re = append(re, []int{})
		}
		re[num] = append(re[num], 1) // 左 1
		if num > 0 {
			last := re[num-1]                  // 获取上一组元素
			for i := 0; i < len(last)-1; i++ { // 相加获得值
				re[num] = append(re[num], last[i]+last[i+1])
			}
			re[num] = append(re[num], 1) // 右 1
		}
		gofxck(num + 1) // 继续递归
	}
	gofxck(0)
	return re
}
