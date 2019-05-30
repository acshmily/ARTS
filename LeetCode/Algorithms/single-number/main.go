package main

import "fmt"

/**
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

func singleNumber(nums []int) int {
	if nums == nil {
		return 0
	}
	a := 0
	for k := range nums {

		a = a ^ nums[k]
	}
	return a
}

func main() {

	fmt.Println(singleNumber([]int{2, 2, 1}))

}
