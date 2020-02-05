package tree

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
//return sum of left child and sum of right child, and tilt
func findTiltHelper(root *TreeNode) (int,  int) {
	if root == nil {
		return 0,  0
	}
	sl, lt := findTiltHelper(root.Left)
	sr, rt := findTiltHelper(root.Right)
	return sl+sr+root.Val, lt+rt+abs(sl, sr)
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return findTiltHelper(root)
}