package main

/**
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its minimum depth = 2.


*/
/**
 * Definition for a binary tree node.         */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	if root.Left == nil && root.Right != nil {
		return r + 1
	}
	if root.Left != nil && root.Right == nil {
		return l + 1
	}
	if l > r {
		return r + 1
	} else {
		return l + 1
	}

}
