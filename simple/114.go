package main

/*
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。



示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
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

// 前序遍历 根左右
func preorderTraversal(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nums
	}

	nums = append(nums, root.Val)
	nums = append(nums, preorderTraversal(root.Left)...)
	nums = append(nums, preorderTraversal(root.Right)...)
	return nums
}
