package main

import (
	"fmt"
	"strings"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

/*
	解题思路：
	1.最优解是可以取这个数组里长度最短的字符串进行遍历（可能需要用到一次for）
	2.然后再循环整个数组判断通用信息
*/

func longestCommonPrefix(strs []string) string {
	commStr := ""
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}
	// 查找长度最小的字符串
	if len(strs[0]) == 0 {
		return ""
	}
	minStr := strs[0]
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < len(minStr) {
			minStr = strs[i]
		}
	}
	for len(minStr) != 0 {
		for i := 0; i < len(strs); i++ {
			if 0 == strings.Index(strs[i], minStr[0:1]) {
				strs[i] = strs[i][1:]
			} else {
				return commStr
			}

		}
		commStr = commStr + minStr[0:1]
		minStr = minStr[1:]
	}
	return commStr

}
func longestCommonPrefix1(strs []string) string {
	// 不存在则返回空白字符串
	if strs == nil {
		return ""
	}
	commonStr := ""
	// 假如只有一个的时候，则返回本身
	if len(strs) == 1 {
		return strs[0]
	} else if len(strs) > 1 { // 当大于1个元素的时候，则需要进行比较操作
		for len(strs[0]) != 0 {
			count := 1
			for i := 1; i < len(strs); i++ {
				if strings.HasPrefix(strs[i], strs[0][0:1]) {
					count++
					strs[i] = strs[i][1:]
				} else {
					return commonStr
				}

			}
			if count == len(strs) {
				commonStr = commonStr + strs[0][0:1]
				strs[0] = strs[0][1:]
			}
		}

	}

	return commonStr
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	str := "abc"
	fmt.Println(str[1:])
}
