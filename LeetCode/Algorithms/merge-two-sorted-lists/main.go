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
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	}
	intList := make([]int, 0, 8)
	re := &ListNode{}
	tmp := re
	for l1.Next != nil || l2.Next != nil {
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
	//判断

	for i := 0; i < len(intList); i++ {
		//第一个ListNode
		tmp.Val = intList[i]
		if i+1 != len(intList) {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}

	}
	return re
}

/**
利用递归
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil { //明确的终止条件
		return l2
	}
	if l2 == nil { //明确的终止条件
		return l1
	}
	var res *ListNode
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
func main() {
	l1 := &ListNode{Val: 5}
	l2 := &ListNode{1, &ListNode{2, &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(l1, l2))
}
