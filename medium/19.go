package main

/*
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
19. 删除链表的倒数第 N 个结点
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。



示例 1：


输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	tmp := head
	for tmp != nil {
		l++
		tmp = tmp.Next
	}

	if n == l {
		return head.Next
	}

	tmp = head
	for i := l - n - 1; i > 0; i-- {
		tmp = tmp.Next
	}

	tmp.Next = tmp.Next.Next
	return head
}
