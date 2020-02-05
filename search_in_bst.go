package tree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == val {
		return root
	}
	//take advantage of the property of BST that left < root < right
	if val > root.Val {
		return searchBST(root.Right, val)
	}
	return searchBST(root.Left, val)
}

//non-recursive solution

func searchBSTNonRecursive(root *TreeNode, val int) *TreeNode {
	for root != nil && root.Val != val {
		if val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return root
}