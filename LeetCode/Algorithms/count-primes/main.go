package main

import "fmt"

/**
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。

*/

func countPrimes(n int) int { // 筛法求质数
	a := make([]bool, n)
	count := 0
	for i := 2; i < n; i++ {
		if a[i] {
			continue
		}
		for j := i + i; j < n; j = j + i {
			a[j] = true
		}
		count++
	}
	return count
}

func main() {
	// 质数只能被自己整除
	fmt.Println(countPrimes(10))

}
