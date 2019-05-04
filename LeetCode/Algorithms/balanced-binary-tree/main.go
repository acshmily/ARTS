package main

import "math"

/**
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:

Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:

Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.


*/
/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	p := depth(root, 1)
	return p != -1
}

func depth(root *TreeNode, high int) int {
	if root == nil { // 如果为空，则表示到达末端节点，返回当前高度
		return high
	} else {
		p1 := depth(root.Left, high+1) // 递归执行左节点
		if p1 == -1 {                  // 当不成立，则返回
			return -1
		}
		p2 := depth(root.Right, high+1) // 递归执行右节点
		if p2 == -1 {                   // 当不成立 则返回
			return -1
		}
		if math.Abs(float64(p1)-float64(p2)) > 1 { // 判断左右两边的高度差不能超过1 否则结束
			return -1
		}
		if p1 > p2 {
			return p1
		} else {
			return p2
		}
	}
}
func main() {

}
