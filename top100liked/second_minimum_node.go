package tree

func min(a, b int) {
	if a > b {
		return b
	}
	return a
}
func findSecondMinimumValue(root *TreeNode) int {
	if root.Left  == nil {
		return -1
	}
	lv, rv := root.Left.Val, root.Right.Val
	lm, rm := lv, rv
	if lv == root.Val {
		lm :=  findSecondMinimumValue(root.Left)
	}
	if rv == root.Val {
		rm := findSecondMinimumValue(root.Right)
	}
	if lm == -1  {
		return rm
	}
	if rm == -1 {
		return lm
	}
	return min(lm, rm)
}
 