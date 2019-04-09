package main

import (
	"fmt"
	"strings"
)

/**
Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.

If the last word does not exist, return 0.

Note: A word is defined as a character sequence consists of non-space characters only.

Example:

Input: "Hello World"
Output: 5
*/
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s) // 去除收尾空格
	if len(s) == 0 {         // 如果就是个空字符串
		return 0
	}
	inx := strings.LastIndex(s, " ")
	if inx == -1 { // 如果没有空格,则返回当前字符串
		return len(s)
	} else {
		return len(s[inx+1:]) // 长度为略过空格 + 1
	}
}
func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}
