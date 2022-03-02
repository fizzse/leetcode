package main

/*

https://leetcode-cn.com/problems/remove-linked-list-elements/
203. 移除链表元素
给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。


示例 1：


输入：head = [1,2,6,3,4,5,6], val = 6
输出：[1,2,3,4,5]
示例 2：

输入：head = [], val = 1
输出：[]
示例 3：

输入：head = [7,7,7,7], val = 7
输出：[]

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	l := newHead
	for l.Next != nil {
		if l.Next.Val == val {
			l.Next = l.Next.Next
			continue
		}

		l = l.Next
	}

	return newHead.Next
}
