package tree

var MAX  = 10000

func min(a, b int) int {
	 if a >= b {
		 return b
	 }
	 return a
}
func abs(a, b int) int {
	if a == MAX || b == MAX {
		return MAX
	}
	if a >= b {
		return a - b
	}
	return b - a
}
func rightMost(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	} 
	return rightMost(root.Right)
}

func leftMost(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return leftMost(root.Left)
}

//the minimum difference from root to any node is 
// min(root - max(left-tree), root-min(right-tree))
func getMinimumDifference(root *TreeNode) int {
	lr := MAX
	rr := MAX
	if root == nil {
		return MAX
	}
	if root.Left != nil {
		lr = abs(root.Val, rightMost(root.Left).Val)
	}
	if root.Right != nil {
		rr = abs(root.Val, leftMost(root.Right).Val)
	}
	//minimum of children
	minimum := min(getMinimumDifference(root.Left),getMinimumDifference(root.Right))

	currentMin := min(lr, rr)
	return min(currentMin, minimum)

}