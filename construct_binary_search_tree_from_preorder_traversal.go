package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstFromPreorder(preorder []int) *TreeNode {
	 if len(preorder) == 0 {
		 return nil
	 }
	 if len(preorder) == 1 {
		 return &TreeNode{Val: preorder[0]}
	 }
	 node := &TreeNode{Val: preorder[0]}
	 left, right := make([]int, 0, 10), make([]int, 0, 10)
	 for _, v := range preorder[1:] {
		 if v < preorder[0] {
			 left = append(left, v)
		 } else {
			 right = append(right, v)
		 }
	 }
	 node.Left = bstFromPreorder(left)
	 node.Right = bstFromPreorder(right)
	 return node
}