package tree

//left is 0, right is 1
func helper(root *TreeNode, b int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && b == 0 {
		return root.Val 
	}
	return helper(root.Left, 0) + helper(root.Right, 1)
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root.Left, 0) + helper(root.Right, 1)
}