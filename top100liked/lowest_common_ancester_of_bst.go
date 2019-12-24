package tree

//since all nodes in the left branch will be smaller than root
// all nodes in the right branch will be larger than the root
// if p and q is smaller than the root, then they are in the left branch
// if p and q is larger than the root, then they are in the right branch
// otherwise the root will be the lowest comman ancester
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pv, qv, rootv := p.Val, q.Val, root.Val
	if pv > rootv && qv > rootv {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if pv < rootv && qv < rootv {
		return lowestCommonAncestor(root.Left, p, q)
	}

	return root
}