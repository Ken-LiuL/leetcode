package tree

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// the login behind root.Left and root.Right part is  that:
	// if root.Left is nil
	//     then if it is not leaf node, we could go minDepth(root.Right) + 1
	//     if it is true leaf node, we could stil go minDepth(root.Right) + 1, since
    //     it would be the same of minDepth(root.Left) + 1
	if root.Left == nil {
		return minDepth(root.Right)+1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	ld, rd := minDepth(root.Left), minDepth(root.Right)
	if ld < rd {
		return ld + 1
	}
	return rd + 1
}
