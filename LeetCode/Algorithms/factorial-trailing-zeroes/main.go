package main

import (
	"fmt"
)

func trailingZeroes(n int) int {
	// 计算5的个数: 5 * 4 * 3 * 2 * 1  中 2 * 5 才会产生 0
	count := 0
	for n >= 5 {
		count += n / 5
		n /= 5
	}
	return count
}
func main() {

	fmt.Println(trailingZeroes(5))

}
