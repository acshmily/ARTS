package main

import "fmt"

/**
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/

/**
Answer:
1.如果是负数可以直接返回false
2.如果为0或者正整数则需要进行数字倒叙排序
3.如果为0结尾的数肯定不会相等

*/
func isPalindrome(x int) bool {
	source := x
	re := 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	for x != 0 {
		pop := x % 10
		x = x / 10
		re = re*10 + pop
	}
	if re == source {
		return true
	} else {
		return false
	}
}
func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(0))

}
