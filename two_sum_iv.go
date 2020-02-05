package tree

func inorder(root *TreeNode, arr *[]int ) {
	if root != nil {
		if root.Left != nil {
			inorder(root.Left, arr)
		}
		*arr = append(*arr, root.Val)
		if root.Right != nil {
			inorder(root.Right, arr)
		}
	}
	return 
	
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	arr := make([]int, 0, 50)
	inorder(root, &arr)
	i, j := 0, len(arr)-1
	for i < j {
		if k - arr[i] > arr[j] {
			i++
		} else if k - arr[i] == arr[j] {
			return true
		}  else {
			j--
		}
	}
	return false
}