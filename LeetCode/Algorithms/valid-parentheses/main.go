package main

import (
	"fmt"
)

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

输入: "()[]{}"
输出: true

输入: "(([]){})"
输出: true
*/

func isValid(s string) bool {
	// 初始化一个map对象储存匹配正确的括号
	valid := map[string]string{"{": "}", "(": ")", "[": "]"} //note!!
	strStack := make([]string, 0, 4)                         // 替换为数组效率很高
	if len(s) == 0 {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	for len(s) != 0 {
		if str, ok := valid[s[0:1]]; ok { // 判断是否为括号开头
			strStack = append(strStack, str)
			s = s[1:] // 去除已经添加的元素
		} else { // 当匹配到关闭符
			if len(strStack) > 0 && s[0:1] == strStack[len(strStack)-1] { // 当stack内有值并且与 需要匹配的括号配对
				s = s[1:]                             // 裁切字符串
				strStack = strStack[:len(strStack)-1] //同样裁切关闭的括号
			} else {
				return false
			}
		}

	}
	if len(strStack) == 0 {
		return true
	} else {
		return false
	}

}
func main() {
	fmt.Println(isValid("(("))
	//s := "12344567"
	//inx := strings.LastIndex(s,"4")
	//s = s[1:inx] + s[inx +1:]
	//fmt.Println(s)
}
