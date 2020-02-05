package tree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left)
	rd := maxDepth(root.Right)
	if ld > rd {
		return ld  + 1
	} else {
		return rd + 1
	}
}