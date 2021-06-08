package main

/*
给定一个二叉树，返回它的 后序 遍历。

示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-postorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 后续遍历 左 右 根
func postorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}

	nums = append(nums, postorderTraversal(root.Left)...)
	nums = append(nums, postorderTraversal(root.Right)...)
	nums = append(nums, root.Val)
	return nums
}

// FIXME 遍历法
