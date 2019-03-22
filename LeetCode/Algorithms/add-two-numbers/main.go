package main

import (
	"fmt"
)

/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
参考数字颠倒方法处理
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//为空的时候
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	prehead := &ListNode{Val: 0} // 申请一个内存块储存信息并获取内存地址
	current := prehead           // 获取链头内存地址
	jw := 0
	for l1 != nil || l2 != nil {
		v1 := 0
		v2 := 0
		if l1 != nil {
			v1 = l1.Val  // 获取当前链的值
			l1 = l1.Next // l1内存地址指向Next地址
		}
		if l2 != nil {
			v2 = l2.Val  // 获取当前链的值
			l2 = l2.Next // l2内存地址指向Next地址
		}
		tmp := v1 + v2 + jw                     // 临时值 = v1 + v2 + jw
		current.Next = &ListNode{Val: tmp % 10} // 指向一个新的内存块 ,第一次执行的时候current 与 prehead 内存地址相同
		current = current.Next
		jw = tmp / 10 //是否进位
		if jw != 0 {
			current.Next = &ListNode{Val: jw}
		}
	}
	return prehead.Next
}

func main() {
	//[1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
	//[5,6,4]
	//l1 := &ListNode{1,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{0,&ListNode{1,nil}}}}}}}}}}}}}}}}}}}}}}}}
	//l2 := &ListNode{5,&ListNode{6,&ListNode{4,nil}}}
	l1 := &ListNode{2, &ListNode{4, &ListNode{Val: 3}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(l1, l2))

}
