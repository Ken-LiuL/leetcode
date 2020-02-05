package tree


func helper(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	return root.Val == val && helper(root.Left, val) && helper(root.Right, val)
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return helper(root, root.Val)
}