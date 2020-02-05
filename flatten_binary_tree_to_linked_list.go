package tree


var pre *TreeNode

func helper(root *TreeNode) {
	if root == nil {
		return
	}
	helper(root.Right)
	helper(root.Left)
	root.Right = pre
	root.Left = nil
	pre = root
}

func flatten(root *TreeNode)  {
	pre = nil
	helper(root)
}