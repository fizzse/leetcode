package main

/*
https://leetcode-cn.com/problems/reverse-linked-list/
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var newHead *ListNode // 指定为nil
	for curr != nil {
		next := curr.Next   // 获取下一个节点
		curr.Next = newHead // 下一个节点指向新头
		newHead = curr      // 新头指定为下一个节点
		curr = next         // 向后移动
	}

	return newHead
}
