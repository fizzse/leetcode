package main

/*
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 转成数组 进行比对
func isPalindrome(head *ListNode) bool {
	var list []int
	node := head
	for node != nil {
		list = append(list, node.Val)
		node = node.Next
	}

	size := len(list)
	j := size - 1
	for i := 0; i < size/2; i++ {
		if list[i] != list[j] {
			return false
		}
		j--
	}

	return true
}
