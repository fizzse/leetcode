package main

import (
	"strconv"
)

/*
给定一个二叉树，返回所有从根节点到叶子节点的路径。

说明: 叶子节点是指没有子节点的节点。

示例:

输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
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
func binaryTreePaths(root *TreeNode) []string {
	pathLeft := ""
	pathRight := ""
	var pathList []string
	if root == nil {
		return pathList
	}

	pathLeft += strconv.Itoa(root.Val)
	pathRight += strconv.Itoa(root.Val)
	if root.Left != nil {
		pathLeft += "->"
		pathLeft += strconv.Itoa(root.Left.Val)
	}

	if pathLeft != "" {
		pathList = append(pathList, pathLeft)
	}

	pathList = append(pathList, binaryTreePaths(root.Left)...)

	if root.Right != nil {
		pathLeft += "->"
		pathLeft += strconv.Itoa(root.Right.Val)
	}

	if pathRight != "" {
		pathList = append(pathList, pathRight)
	}

	pathList = append(pathList, binaryTreePaths(root.Right)...)
	return pathList
}

func onePath(root *TreeNode, path string) string {
	if root == nil {
		return path
	}

	path += strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return path
	}

	onePath(root.Left, path)
	onePath(root.Right, path)
	return path
}
