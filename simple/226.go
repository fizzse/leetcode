package main

/*
翻转一棵二叉树。

示例：

输入：

     4
   /   \
  2     7
 / \   / \
1   3 6   9
输出：

     4
   /   \
  7     2
 / \   / \
9   6 3   1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/invert-binary-tree
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

//func invertTree(root *TreeNode) *TreeNode {
//	/*
//	 对称二叉树的思路
//	*/
//
//	if root == nil {
//		return nil
//	}
//
//	swap(root.Left, root.Right)
//	return root
//}
//
//func swap(r1, r2 *TreeNode) {
//	// 都不为空
//	if r1 == nil || r2 == nil {
//		r1, r2 = r2, r1 // 这样交换 对root的引用关系并没有影响
//		return
//	}
//
//	r1.Val, r2.Val = r2.Val, r1.Val
//	swap(r1.Left, r2.Right)
//	swap(r1.Right, r2.Left)
//}

func invertTree(root *TreeNode) *TreeNode {
	// 遍历整棵树 将左右互换即可达到 翻转的目的

	if root == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
