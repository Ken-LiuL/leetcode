package   tree

//since the right branch will always larger than itself, so we convert from right,
//and return all the sum of right and add it to self's value
var sum int
func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	convertHelper(root)
	return root
}

func convertHelper(root *TreeNode)  {
	if root == nil {
		return 
	}
    convertHelper(root.Right)
	root.Val += sum
    sum = root.Val
	convertHelper(root.Left)
 }

