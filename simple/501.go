package main

import "fmt"

/*
501. 二叉搜索树中的众数
给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。

如果树中有不止一个众数，可以按 任意顺序 返回。

假定 BST 满足如下定义：

结点左子树中所含节点的值 小于等于 当前节点的值
结点右子树中所含节点的值 大于等于 当前节点的值
左子树和右子树都是二叉搜索树
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var res []int
	nums := getNums(root)

	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	maxCnt := 0
	for _, v := range m {
		if v > maxCnt {
			maxCnt = v
		}
	}

	for k, v := range m {
		if v == maxCnt {
			res = append(res, k)
		}
	}

	return res
}

func getNums(root *TreeNode) []int {
	var nums []int
	if root == nil {
		return nil
	}

	nums = append(nums, root.Val)
	if root.Left != nil {
		nums = append(nums, getNums(root.Left)...)
	}
	if root.Right != nil {
		nums = append(nums, getNums(root.Right)...)
	}

	return nums
}

func main() {
	n := &TreeNode{Val: 0}
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 2}
	n.Left = n1
	n.Right = n2
	n1.Left = n3

	o := findMode(n)
	fmt.Println(o)
}
