package tree

 type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
 }
 func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p != nil && q != nil && p.Val == q.Val {
		return  isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else if p == nil && q == nil {
		return true
	}
	else {
		return false
	}
}