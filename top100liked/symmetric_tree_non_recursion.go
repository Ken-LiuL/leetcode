package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(tree1 *TreeNode, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	} else if tree1 == nil || tree2 == nil {
		return false
	}
	return tree1.Val == tree2.Val && isMirror(tree1.Left, tree2.Right) && isMirror(tree1.Right, tree2.Left)
}