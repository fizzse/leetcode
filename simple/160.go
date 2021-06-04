package main

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。

图示两个链表在节点 c1 开始相交：

题目数据 保证 整个链式结构中不存在环。

注意，函数返回结果后，链表必须 保持其原始结构 。

示例 1：


输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Intersected at '8'
解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
示例 2：


输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
输出：Intersected at '2'
解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。
示例 3：



输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
输出：null
解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
这两个链表不相交，因此返回 null 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-linked-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// map 偷鸡
/*func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)

	for headA != nil {
		m[headA]++
		headA = headA.Next
	}

	for headB != nil {
		m[headB]++
		if m[headB] == 2 {
			return headB
		}
		headB = headB.Next
	}

	return nil
}*/

// 链表有环的变种
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	/*
	 链表遍历到尾部 就从另外一条链表的开头开始遍历
	 如果有交点 总会相遇的。没有交点 最后 nil == nil 退出循环
	 如果长度相等 第一轮就能得到结果 如果长度不相等 总会有相遇的时候
	*/

	if headA == nil || headB == nil {
		return nil
	}

	h1, h2 := headA, headB
	for h1 != h2 {
		if h1.Next == nil {
			h1 = headB
		} else {
			h1 = h1.Next
		}

		if h2.Next == nil {
			h2 = headA
		} else {
			h2 = h2.Next
		}
	}

	return h1
}