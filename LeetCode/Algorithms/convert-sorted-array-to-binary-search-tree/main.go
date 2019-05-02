package main

/**
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted array: [-10,-3,0,5,9],

One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
*/

/**
 * Definition for a binary tree node.**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil {
		return nil
	} else {
		return returnBST(nums, 0, len(nums)-1)
	}

}
func returnBST(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	m := l + (r-l)/2
	root := &TreeNode{}
	root.Val = nums[m]
	root.Left = returnBST(nums, l, m-1)
	root.Right = returnBST(nums, m+1, r)
	return root
}
func main() {

}
