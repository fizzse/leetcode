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
//func isPalindrome(head *ListNode) bool {
//	node := head
//	m := make(map[int]int)
//	i := 0
//	for node != nil {
//		m[node.Val] += i
//		node = node.Next
//		i++
//	}
//
//	mid := i / 2
//	fmt.Println(mid)
//	for _, v := range m {
//		if v%mid != 0 {
//			return false
//		}
//	}
//
//	return true
//}
