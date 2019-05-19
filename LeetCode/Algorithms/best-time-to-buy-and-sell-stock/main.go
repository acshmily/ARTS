package main

/**
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/
// 当 第1天 买 7 第2天 卖  利润为 -6
//				第3天 卖	 利润为 -2
//				第4天 卖  利润为 -4
//				5			   -1
//				6			   -3
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	min, max := prices[0], 0
	for i := 0; i < len(prices); i++ {
		max = maxInt(max, prices[i]-min)
		min = minInt(min, prices[i])
	}
	return max
}
func maxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}

}
func minInt(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	println(maxProfit([]int{1, 2}))
}
