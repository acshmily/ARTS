package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/
//Answer1 : 字符串转换
func reverse_str(x int) int {

	str := strconv.Itoa(x)
	var re string = ""
	// 如果为正整数
	if x > 0 {
		intLen := len(strconv.Itoa(x)) //转为str
		for i := 0; i < intLen; i++ {

			re = str[i:i+1] + re

		}
	} else if x < 0 {

		str = str[1:]      // 截取负号
		intLen := len(str) //转为str
		for i := 0; i < intLen; i++ {

			re = str[i:i+1] + re

		}
		re = "-" + re
	} else {
		re = "0"
	}
	nint, error := strconv.Atoi(re)
	if error != nil || nint > math.MaxInt32 || nint < math.MinInt32 {
		return 0
	}

	return nint
}

// Answer2 取模运算
func reverse(x int) int {
	re := 0
	for x != 0 {
		pop := x % 10
		x = x / 10
		re = re*10 + pop
	}
	if re > math.MaxInt32 || re < math.MinInt32 {
		return 0
	}
	return re
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(123000))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(123456789))
}
