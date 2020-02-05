package tree


func invertTree(root *TreeNode) *TreeNode {
	 root.Left, root.Right = root.Right, root.Left
	 invertTree(root.Left)
	 invertTree(root.Right)
}