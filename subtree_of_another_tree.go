package tree

func checkExactSame(s *TreeNode, t *TreeNode) bool {
	if t == nil || s == nil {
		return  t == s
	}
	return s.Val == t.Val  && checkExactSame(s.Left, t.Left) && checkExactSame(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil || s == nil {
		return  t == s
	}
	return checkExactSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}