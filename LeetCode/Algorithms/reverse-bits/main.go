package main

import "fmt"

/**
输入: 00000010100101000001111010011100
输出: 00111001011110000010100101000000
解释: 输入的二进制串 00000010100101000001111010011100 表示无符号整数 43261596，
      因此返回 964176192，其二进制表示形式为 00111001011110000010100101000000。
*/
func reverseBits(num uint32) uint32 {
	var temp uint32 = 0
	length := 32
	for length > 0 {
		temp <<= 1      //左移后赋值	C <<= 2 等于 C = C << 2
		temp |= num & 1 //按位或后赋值	C |= 2 等于 C = C | 2
		num >>= 1       //	右移后赋值C >>= 2 等于 C = C >> 2
		length--
	}
	return temp
}

func main() {
	fmt.Printf("%s", reverseBits(00000010100101000001111010011100))
}
