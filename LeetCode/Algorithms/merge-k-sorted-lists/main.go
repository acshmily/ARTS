package main

import "sort"

/**
 * Definition for singly-linked list.
	Merge k sorted linked lists and return it as one sorted list.
	Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力分解法
// todo 未考虑时间复杂度
func mergeKLists(lists []*ListNode) *ListNode {
	reList := &ListNode{}
	numlist := make([]int, 0)
	for i := 0; i < len(lists); i++ { // 循环每个链表
		list := lists[i]
		for list != nil { // 循环每个链表元素
			numlist = append(numlist, list.Val)
			list = list.Next
		}
	}
	sort.Ints(numlist)
	x := reList
	for _, e := range numlist {
		tmp := &ListNode{}
		tmp.Val = e
		x.Next = tmp
		x = x.Next

	}
	return reList.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	mergeKLists([]*ListNode{l1, l2})

}
