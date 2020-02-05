package tree

func helper(root *TreeNode, prefix int) int { 
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return prefix << 1 + root.Val
	}
	return helper(root.Left, newPrefix) + helper(root.Right, newPrefix)
}

func sumRootToLeaf(root *TreeNode) int {
	return helper(root, 0) 
}