package main

var ans = 0

func findTilt(root *TreeNode) int {
	calu(root)
	return ans
}

func calu(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := calu(root.Left)
	right := calu(root.Right)

	ans += absInt(left - right)
	return root.Val + left + right
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
