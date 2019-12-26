package tree


func helper(root *TreeNode, arr  *[]int) {
	if root ==  nil {
		return 
	}
	helper(root.Left, arr)
	*arr = append(*arr, element)
	helper(root.Right, arr)
}

func inorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0, 20)
	return helper(root, &arr)
}
