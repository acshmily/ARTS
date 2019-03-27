package main

import (
	"fmt"
	"sort"
)

/**
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 注意为有序
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	intList := make([]int, 0, 8)
	re := &ListNode{}
	tmp := re
	for l1.Next != nil && l2.Next != nil {
		if l1.Next != nil {
			intList = append(intList, l1.Val)
			l1 = l1.Next
		}
		if l2.Next != nil {
			intList = append(intList, l2.Val)
			l2 = l2.Next
		}
	}
	intList = append(intList, l1.Val, l2.Val)
	sort.Ints(intList) // 排序
	for i := 0; i < len(intList); i++ {
		tmp = &ListNode{Val: intList[i]}
		tmp.Next = &ListNode{}
		re.Next = tmp

	}
	return re.Next
}
func main() {
	l1 := &ListNode{1, &ListNode{2, &ListNode{Val: 4}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(l1, l2))
}
