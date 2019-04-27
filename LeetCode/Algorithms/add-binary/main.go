package main

import "fmt"

/**
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/
func addBinary(a string, b string) string {
	aChar, bChar := []rune(a), []rune(b)
	var s []rune
	var c rune
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || c == 1 {
		if i >= 0 {
			c += aChar[i] - '0'
			i--
		}
		if j >= 0 {
			c += bChar[j] - '0'
			j--
		}
		s = append(s, c%2+'0')
		c /= 2
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func addBinary2(a string, b string) string {
	s := make([]byte, 0)
	var c byte
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 || c == 1 {
		if i >= 0 {
			c += a[i] - '0'
			i--
		}
		if j >= 0 {
			c += b[j] - '0'
			j--
		}
		s = append(s, c%2+'0') // 转byte\
		c /= 2
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)

}

func main() {
	s := byte('1')
	fmt.Print(s)
}
