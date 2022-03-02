package main

import (
	"container/list"
	"fmt"
)

/*
232. 用栈实现队列
请你仅使用两个栈实现先入先出队列。队列应当支持一般队列支持的所有操作（push、pop、peek、empty）：

实现 MyQueue 类：

void push(int x) 将元素 x 推到队列的末尾
int pop() 从队列的开头移除并返回元素
int peek() 返回队列开头的元素
boolean empty() 如果队列为空，返回 true ；否则，返回 false
说明：

你 只能 使用标准的栈操作 —— 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
*/

type MyQueue struct {
	s1 *myStack // 写入栈
	s2 *myStack // pop栈
}

func Constructor() MyQueue {
	return MyQueue{
		s1: &myStack{},
		s2: &myStack{},
	}
}

func (this *MyQueue) Push(x int) {
	this.s1.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.s2.Len() > 0 {
		return this.s2.Pop()
	}
	for this.s1.Len() > 0 {
		n := this.s1.Pop()
		this.s2.Push(n)
	}

	return this.s2.Pop()
}

func (this *MyQueue) Peek() int {
	if this.s2.Len() > 0 {
		return this.s2.Peek()
	}

	for this.s1.Len() > 0 {
		n := this.s1.Pop()
		this.s2.Push(n)
	}
	return this.s2.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.s1.Len() == 0 && this.s2.Len() == 0
}

//--------------------------------------------------实现一个栈--------------------------------------------------------//

type myStack struct {
	l1 list.List
}

func (this *myStack) Push(x int) {
	this.l1.PushFront(x)
}

func (this *myStack) Pop() int {
	if this.l1.Len() == 0 {
		return -1
	}

	head := this.l1.Front()
	this.l1.Remove(head)
	return head.Value.(int)
}

func (this *myStack) Peek() int {
	if this.l1.Len() == 0 {
		return -1
	}

	head := this.l1.Front()
	return head.Value.(int)
}

func (this *myStack) Len() int {
	return this.l1.Len()
}

func (this *myStack) Print() {
	if this.l1.Len() == 0 {
		return
	}

	head := this.l1.Front()
	for i := 0; i < this.l1.Len(); i++ {
		fmt.Printf("%v\t", head.Value)
		head = head.Next()
	}
	fmt.Println()
}

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	n := q.Pop()
	fmt.Println(n)
	q.Push(3)
	q.Push(4)
	n = q.Pop()
	fmt.Println(n)
	n = q.Pop()
	fmt.Println(n)
	n = q.Pop()
	fmt.Println(n)
}
