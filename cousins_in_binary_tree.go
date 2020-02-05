package tree

func helper(p *TreeNode, root *TreeNode, val int, depth int) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	} 
    if root.Val == val {
		return  p, depth
    }
	lp, ld := helper(root, root.Left, val, depth+1)
 	if lp == nil {
		return helper(root, root.Right, val, depth+1)
	} else {
		return lp, ld
	}
}

func isCousins(root *TreeNode, x int, y int) bool {
		lp, ld := helper(root, root, x, 1)
		rp, rd := helper(root, root, y, 1)
		return lp != rp && ld == rd
}