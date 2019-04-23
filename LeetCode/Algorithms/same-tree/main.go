package main

import "fmt"

/**
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

Example 1:

Input:     1         1
          / \       / \
         2   3     2   3

        [1,2,3],   [1,2,3]

Output: true
Example 2:

Input:     1         1
          /           \
         2             2

        [1,2],     [1,null,2]

Output: false
Example 3:

Input:     1         1
          / \       / \
         2   1     1   2

        [1,2,1],   [1,1,2]

Output: false
*/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ints1 = make([]int, 0)
var ints2 = make([]int, 0)

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil && p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
	return true
}

//先序查找
func preOrder(node *TreeNode, ints []int) {
	if node != nil {
		// 当前节点的根
		ints1 = append(ints, node.Val)
		// 查找左节点
		preOrder(node.Left, ints1)
		// 查找又节点
		preOrder(node.Right, ints1)
	}
}
func main() {
	node := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: nil, Left: nil}, Right: &TreeNode{Val: 3, Right: nil, Left: nil}}
	preOrder(node, ints1)
	fmt.Println(ints1)
}
