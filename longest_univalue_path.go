package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var longest = 0

func longestPathHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := longestPathHelper(root.Left), longestPathHelper(root.Right)

	if root.Left != nil && root.Left.Val == root.Val {
		left += 1
	} else {
		left = 0
	}
	if root.Right != nil && root.Right.Val == root.Val {
		right += 1
	} else {
		right = 0
	}

	longest = max(longest, left+right)
	return max(left, right)
}

func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	longestPathHelper(root)
	res := longest
	longest = 0
	return res
}
