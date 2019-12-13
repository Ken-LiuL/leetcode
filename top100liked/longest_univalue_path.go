package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPathHelper(root *TreeNode, l int) int {
	cur := l
	if root.Left != nil {
		if root.Val == root.Left.Val {
			cur = max(cur, longestPathHelper(root.Left, l+1))
		} else {
			cur = max(cur, longestPathHelper(root.Left, l))
		}
	}
	if root.Right != nil {
		if root.Val == root.Right.Val {
			cur = max(cur, longestPathHelper(root.Right, l+1))
		} else {
			cur = max(cur, longestPathHelper(root.Right, l))
		}
	}
	return cur
}

func longestUnivaluePath(root *TreeNode) int {
	return longestPathHelper(root, 0)
}
