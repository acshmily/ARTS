package main

import "fmt"

/**
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().
*/
func strStr(haystack string, needle string) int {
	inx := -1
	if len(needle) == 0 {
		return 0
	}
	if len(needle) >= len(haystack) && haystack != needle {
		return -1
	}
	for i := 0; i < len(haystack); i++ {
		if haystack[i] != needle[0] {
			continue
		}
		if len(haystack[i:]) < len(needle) {
			break
		}
		if haystack[i:i+len(needle)] == needle {
			inx = i
			break
		}

	}
	return inx

}
func main() {
	fmt.Println(strStr("aaaaa", "bba"))
}
