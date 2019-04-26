package main

/**
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
Note:
Bonus points if you could solve it both recursively and iteratively.
*/
/**
 * Definition for a binary tree node.
 *
**/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkmirr(root.Left, root.Right)

}
func checkmirr(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil { // 出口：如果左树节点和又树节点都为空，则表示已经到末节点了
		return true
	}
	if a == nil || b == nil { // 出口：如果左节点或右节点有一个为空
		return false
	}
	if a.Val != b.Val { // 出口： 值不相等
		return false
	} else {
		return checkmirr(a.Left, b.Right) && checkmirr(a.Right, b.Left)
	}
}
func main() {
	tn := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2, Left: nil, Right: &TreeNode{Val: 3}}}
	isSymmetric(tn)
}
