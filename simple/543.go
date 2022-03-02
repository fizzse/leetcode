package main

/*
543. 二叉树的直径
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。



示例 :
给定二叉树

          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者 [5,2,1,3]。



注意：两结点之间的路径长度是以它们之间边的数目表示。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 等于左子树的深度 + 右子树的深度

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(depth(node.Left), depth(node.Right)) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(diameterOfBinaryTree(root.Right), max(diameterOfBinaryTree(root.Left), depth(root.Left)+depth(root.Right)))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
