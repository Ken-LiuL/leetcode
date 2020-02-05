package tree


func hasPathSum(root *TreeNode, sum int) bool {
	if root != nil && root.Left == nil && root.Right == nil{
		if sum == root.Val {
			return true
		} else {
			return false
		}
	}
	if root == nil {
		return false
	}
	return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}