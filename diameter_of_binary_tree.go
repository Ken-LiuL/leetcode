package tree


var max int

//max  =  l + r, max
//the longest path of current tree is max(l, r) + 1
func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	helper(root)
	return max
}

func maxFn(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(root *TreeNode) int { 
	if root == nil {
		return 0
	}
	l, r := helper(root.Left), helper(root.Right)
	max =  maxFn(max, l+r)
	return maxFn(l, r ) + 1

}


