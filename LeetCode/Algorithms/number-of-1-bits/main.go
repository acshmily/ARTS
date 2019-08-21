package main

/**
输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

*/

func hammingWeight(num uint32) int {
	count := 0     // 定义统计变量
	for num != 0 { //位运算操作
		if num&1 == 1 { // 与 1 与操作 如果位1 则位1
			count++
		}
		num >>= 1 // 右移1 位
	}
	return count
}
