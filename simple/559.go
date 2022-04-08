package main

/*
https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
559. N 叉树的最大深度
给定一个 N 叉树，找到其最大深度。

最大深度是指从根节点到最远叶子节点的最长路径上的节点总数。

N 叉树输入按层序遍历序列化表示，每组子节点由空值分隔（请参见示例）。
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	for _, c := range root.Children {
		depth = getMax(depth, maxDepth(c))
	}

	return depth + 1
}

func getMax(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}
