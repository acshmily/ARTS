package main

import (
	"fmt"
	"math"
)

/**
最长回文子串
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/
/**
快慢指针问题
*/

func longestPalindrome(s string) string {
	if len(s) < 2 { // 肯定是回文，直接返回
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)                //a [b] a bd 可能回文数的中间数是一个 则从一个数开始两边扩散
		len2 := expandAroundCenter(s, i, i+1)              //a [bb] c 可能回文数的中间数可以使两个。从两个开始两边是扩散
		len := int(math.Max(float64(len1), float64(len2))) //比较出两个的最长的长度
		if len > end-start {
			//len=3  i=1
			start = i - (len-1)/2 //i:在这里就是一个回文数的中间数，（len-1）:是处理len可能是偶数的情况
			end = i + len/2
		}
	}
	return string(s[start : end+1])

}

/**
展开查询
*/
func expandAroundCenter(s string, left, right int) int {
	L, R := left, right
	for L >= 0 && R < len(s) && s[L] == s[R] {
		L--
		R++
	}
	return R - L - 1
}
func main() {
	fmt.Println(longestPalindrome("ac"))

}
