package main

/*
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
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
func sumOfLeftLeaves(root *TreeNode) int {
	num := 0
	if root == nil {
		return num
	}

	//if root.Left != nil { // 这样不一定是叶子节点
	//	num += root.Left.Val
	//}

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil { // 确保是叶子
			num += root.Left.Val
		}
	}

	num += sumOfLeftLeaves(root.Left)
	num += sumOfLeftLeaves(root.Right)
	return num
}
