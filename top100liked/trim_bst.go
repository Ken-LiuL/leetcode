package tree



func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}
	//if current value is smaller than the L bound, then all nodes in left
	//branch will be smaller than L, so we drup thte left branch
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}
	//if current value is larger than the R bound, then all nodes in the right branch
	//will be larger than R, so we drop the right branch
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	//recursively walk down
	root.Left = trimBST(root.Left)
	root.Right = trimBST(root.Right)
	return root
}