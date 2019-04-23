package main

import (
	"fmt"
	"math"
)

/**
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: 3
Output: "III"
Example 2:

Input: 4
Output: "IV"
Example 3:

Input: 9
Output: "IX"
Example 4:

Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

/**
规则要求 最大为  3999
*/
func intToRoman(num int) string {
	dict := [4][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 个位上所有可能的值
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 十位上所有可能的值
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 百位上所有可能的值
		{"", "M", "MM", "MMM"}, // 千位上所有可能的值
	}

	return dict[3][num/1000] + dict[2][num/100%10] + dict[1][num/10%10] + dict[0][num%10]

}

var roman = map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M", 4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}

// 暴力破解法
func intToRoman1(num int) string {
	// 将所有情况存入map
	restr := ""
	carry := 0
	if val, ok := roman[num]; ok {
		return val
	}
	// 查询每个数字
	for num != 0 {
		remainder := num % 10                                       // 取的余数
		num = num / 10                                              // 取的商
		if val, ok := roman[remainder*int(math.Pow10(carry))]; ok { // 如果存在特殊
			restr = val + restr
		} else { // 不存在则判断
			if remainder < 4 {
				for remainder != 0 {
					restr = roman[int(math.Pow10(carry))] + restr
					remainder--
				}
			} else if remainder > 5 {
				remainder = remainder - 5
				for remainder != 0 {
					restr = roman[int(math.Pow10(carry))] + restr
					remainder--
				}
				restr = roman[5*int(math.Pow10(carry))] + restr
			}
		}
		carry++
	}
	return restr
}
func main() {
	fmt.Println(intToRoman(3))
}
