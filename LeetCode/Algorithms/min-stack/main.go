package main

import "math"

/**
/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
*/
/**
栈: 先进后出
push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。
*/
type MinStack struct {
	val  int       //定义当前值
	min  int       //定义最小值
	next *MinStack //链表下一指针
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{0, 0, nil}
}

/**
push(x) -- 将元素 x 推入栈中。
*/
func (this *MinStack) Push(x int) {
	min := x
	// 比较最小值
	if this.next != nil {
		min = int(math.Min(float64(x), float64(this.next.min)))
	}
	// 创建一个新的node,并且next指向原来的node 完成压栈动作
	stack := &MinStack{x, min, this.next}
	this.next = stack

}

/**
pop() -- 删除栈顶的元素。
*/
func (this *MinStack) Pop() {
	if this.next == nil {
		return
	}
	this.next = this.next.next
}

/**
top() -- 获取栈顶元素。
*/
func (this *MinStack) Top() int {
	if this.next == nil {
		return 0
	}
	return this.next.val
}

func (this *MinStack) GetMin() int {
	if this.next == nil {
		return 0
	}
	return this.next.min
}
