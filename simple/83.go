package main

/*
存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。

返回同样按升序排列的结果链表。



示例 1：


输入：head = [1,1,2]
输出：[1,2]
示例 2：


输入：head = [1,1,2,3,3]
输出：[1,2,3]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*type ListNode struct {
	Val  int
	Next *ListNode
}*/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	pre := head
	cur := head.Next
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}

		cur = cur.Next
		pre = pre.Next
	}

	return head
}
