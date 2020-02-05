package recursion

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var MAX_V = math.MaxInt32

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxLeft(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val

}
func minRight(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left.Val
	}
	return root.Val

}

/*
*  for any node n , the possible node with minimum distance woule be  the left right node
*  and right left node
*/

func minDiffHelper() {
	m := MAX_V
	if root.Left != nil {
		m = min(root.Val-maxLeft(root.Left), minDiffInBST(root.Left))
	}
	if root.Right != nil {
		m = min(m, minRight(root.Right)-root.Val)
		m = min(m, minDiffInBST(root.Right))
	}
	return m
}
