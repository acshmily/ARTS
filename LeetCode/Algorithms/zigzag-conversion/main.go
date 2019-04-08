package main

import "math"

/**
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	maxLength := math.Min(float64(numRows), float64(len(s)))
	bytes := make([][]byte, int(maxLength))
	//设置控制移动的标志变量
	goingDownFlag := false
	currRow := 0
	for _, val := range s {
		bytes[currRow] = append(bytes[currRow], byte(val))
		if currRow == 0 || currRow == numRows-1 { // currRow == 0  为开头  currRow == numRows - 1 为到底
			goingDownFlag = !goingDownFlag
		}
		if goingDownFlag { // 如果没到底
			currRow++
		} else { // 到底了
			currRow--
		}
	}
	//遍历并输出最后的结果
	result := ""
	for _, v := range bytes {
		result += string(v)
	}
	return result

}

func main() {

}
