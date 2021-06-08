package main

/*
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。



示例 1：


输入：root = [3,9,20,null,null,15,7]
输出：true
示例 2：


输入：root = [1,2,2,3,3,null,null,4,4]
输出：false
示例 3：

输入：root = []
输出：true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/balanced-binary-tree
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
func isBalanced(root *TreeNode) bool {
	/*
	 左子树与右子树的差值不能大于1
	*/
	return deep(root) >= 0
}

// 最大深度
func deep(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deepLeft := deep(root.Left)
	deepRight := deep(root.Right)

	// 不是平衡状态 直接return
	if deepLeft == -1 || deepRight == -1 {
		return -1
	}

	// 比较两棵树的差值
	diff := deepLeft - deepRight
	if diff >= -1 && diff <= 1 {
		if diff >= 0 {
			return deepLeft + 1
		}

		return deepRight + 1
	}

	return -1
}
