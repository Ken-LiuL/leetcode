package tree

var pre int
var max int
var occurence int
var res []int

func findMaxOccurence(root *TreeNode) {
	if pre == root.Val {
		occurence++
	} else {
		occurence = 1
	}
	if occurence > max {
		max = occurence
	}
}

func findAllElements(root *TreeNode) {
	if pre == root.Val {
		occurence++
	} else {
		occurence = 1
	}
	if occurence == max {
		res = append(res, root.Val)
		occurence = 0
	}
}
func inorder(root *TreeNode, f func(*TreeNode)) {
	if root == nil {
		return 
	}
	inorderj(root.Left, f)
	f(root)
	pre = root.Val
	inorder(root.Right, f)
}
func findMode(root *TreeNode) []int {
	//find max occurence first by using inorder recursion
	pre, max = 0, 0 
	inorder(root, findMaxOccurence)
	pre, occurence = 0, 0
	res = make([]int, 0, 10)
	//find all elements with max occurence
	inorder(root, findAllElements)
	return res
}
