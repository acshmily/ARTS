package main

/**
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its bottom-up level order traversal as:
[
  [15,7],
  [9,20],
  [3]
]
*/
/**
 * Definition for a binary tree node.
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	re := [][]int{}
	var dfs func(*TreeNode, int)            // 定义方法
	dfs = func(root *TreeNode, level int) { // 定义方法体
		if root == nil { // 当已经为空则结束
			return
		}
		if level >= len(re) { // 当有新的层
			re = append([][]int{}, re...) // 增加容量
		}
		n := len(re)
		re[n-level-1] = append(re[n-level-1], root.Val)

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}

	dfs(root, 0)

	return re

}
