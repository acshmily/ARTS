package main

import (
	"fmt"
)

/**
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return nil
	}
	target := true
	for i := len(digits) - 1; i >= 0; i-- {
		if target {
			digits[i]++ //如果已经需要进位了
			if digits[i] == 10 {
				digits[i] = 0
				target = true
				if i == 0 {
					tmp := []int{1}
					digits = append(tmp, digits...)
				}
				continue
			}
			target = false
		}

	}
	return digits
}
func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))

}
