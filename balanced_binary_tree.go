package tree

//-1 represent it is not a balanced tree
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	hl := getHeight(root.Left)
	if hl == -1 {
		return -1
	}
	rl := getHeight(root.Right)
	if rl == -1 {
		return -1
	}

	diff := hl  - rl
	if diff > 1 || diff < -1  {
		return -1
	}
	if hl > rl {
		return hl + 1
	}
	return rl + 1
}

func isBalanced(root *TreeNode) bool {
	if getHeight(root) == -1 {
		return false
	}
	return true
}