package tree

func helper(root *TreeNode, arr *[]int) {
	if root == nil {
		return 
	}
	helper(root.Left, arr)
	helper(root.Right, arr)
	*arr = append(*arr, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0, 20)
	helper(root, &arr)
	return arr
}