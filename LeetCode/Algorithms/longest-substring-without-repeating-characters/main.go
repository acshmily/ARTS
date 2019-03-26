package main

import "fmt"

/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
/**
用空间复杂度换时间复杂度
*/
func lengthOfLongestSubstring(s string) int {
	maxLength := 0 //最大长度
	start := 0
	strMap := make(map[string]int)
	sLength := len(s)
	// 循环每个字符串
	for i := 0; i < sLength; i++ {
		if inx, ok := strMap[s[:1]]; ok && start <= inx { // 如果已经存在 这里start <= inx
			start = inx + 1
		}
		if i-start+1 > maxLength { // 当前的非重复字符串长度是否大于 之前的最大非重复长度
			maxLength = i - start + 1
		}
		strMap[s[:1]] = i
		s = s[1:]
	}
	//fmt.Println(strMap)
	return maxLength
}
func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("abac"))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
