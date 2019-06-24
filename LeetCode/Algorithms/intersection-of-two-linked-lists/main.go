package main

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
	定义两个指针, 第一轮让两个到达末尾的节点指向另一个链表的头部,
	最后如果相遇则为交点(在第一轮移动中恰好抹除了长度差)
	两个指针等于移动了相同的距离, 有交点就返回, 无交点就是各走了两条指针的长度
**/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA := headA
	pB := headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
