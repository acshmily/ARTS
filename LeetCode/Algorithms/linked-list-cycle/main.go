package main

/**
 * Definition for singly-linked list.
**/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	root := &ListNode{}
	p := head
	for p != nil {
		q := p.Next
		if p == root {
			return true
		}
		p.Next = root
		p = q
	}
	return false
}
