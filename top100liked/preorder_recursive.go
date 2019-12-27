package tree


func helper(root *TreeNode, arr *[]int)  {
	if root == nil {
		return
	}
	*arr = append(*arr, root.Val)
	helper(root.Left, arr)
	helper(root.Right, arr)
}

func preorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0, 10)
	return helper(root, &arr)
}