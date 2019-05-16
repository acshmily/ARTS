package main

/**
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

Note that the row index starts from 0.


In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:

Input: 3
Output: [1,3,3,1]
Follow up:
*/

func getRow(rowIndex int) []int {
	re := make([]int, rowIndex+1) // 构造一个一维数组
	re[0] = 1                     // 定义边
	for i := 1; i <= rowIndex; i++ {
		for j := i; j >= 1; j-- {
			re[j] = re[j] + re[j-1]
		}
	}
	return re
}
