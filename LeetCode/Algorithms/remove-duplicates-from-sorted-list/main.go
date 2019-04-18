package main

/**
 * Definition for singly-linked list.
	Given a sorted linked list, delete all duplicates such that each element appear only once.

	Example 1:

	Input: 1->1->2
	Output: 1->2
	Example 2:

	Input: 1->1->2->3->3
	Output: 1->2->3
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == tmp.Val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head

}

func main() {

}
