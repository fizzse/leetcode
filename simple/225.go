package main

import "container/list"

/*
https://leetcode-cn.com/problems/implement-stack-using-queues/
225. 用队列实现栈
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。


注意：

你只能使用队列的基本操作 —— 也就是 push to back、peek/pop from front、size 和 is empty 这些操作。
你所使用的语言也许不支持队列。 你可以使用 list （列表）或者 deque（双端队列）来模拟一个队列 , 只要是标准的队列操作即可。

*/

type MyStack struct {
	q1 *myQueue
	q2 *myQueue
}

func Constructor() MyStack {
	return MyStack{
		q1: &myQueue{},
		q2: &myQueue{},
	}
}

func (this *MyStack) Push(x int) { // 在不为空的队列中插入
	q1 := this.q1
	if this.q2.Len() > 0 {
		q1 = this.q2
	}
	q1.Push(x)
}

func (this *MyStack) Pop() int { // 想空队列转移 pop最后一个
	q1 := this.q1
	q2 := this.q2
	if q2.Len() > 0 {
		q1, q2 = q2, q1
	}

	for q1.Len() > 1 {
		n := q1.Pop()
		q2.Push(n)
	}

	return q1.Pop()
}

func (this *MyStack) Top() int {
	n := this.Pop()
	this.Push(n)
	return n
}

func (this *MyStack) Empty() bool {
	return this.q1.Len() == 0 && this.q1.Len() == 0
}

//-----------------------------------queue-------------------------------//

type myQueue struct {
	l1 list.List
}

func (this *myQueue) Push(x int) {
	this.l1.PushBack(x)
}

func (this *myQueue) Pop() int {
	head := this.l1.Front()
	this.l1.Remove(head)
	return head.Value.(int)
}

func (this *myQueue) Len() int {
	return this.l1.Len()
}

func (this *myQueue) Top() int {
	head := this.l1.Front()
	return head.Value.(int)
}
