package tree

type queue []*TreeNode

func newQueue() *queue {
	return &queue{}
}

func (q *queue) push(n *TreeNode) {
	  *q = append(*q, n)
}

func (q *queue) pop() *TreeNode{
	 tmp := (*q)[0]
	 *q = (*q)[1:]
	 return tmp
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func inorder(root *TreeNode, q *queue, check bool) bool {
	 res := true
	 if root == nil {
		 return res
	 }
	 if root.Left == nil &&  root.Right == nil {
	   if !check {
		 q.push(root)
		 res = true
	 	} else {
		 res = q.pop().Val == root.Val
	 	}
	 } 
	 if root.Left != nil {
		res = inorder(root.Left, q, check)
	 } 
	 if root.Right != nil {
		 res = res && inorder(root.Right, q, check)
	 }
	 return res
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	 q := newQueue()
	 inorder(root1, q, false) 
	 return inorder(root2, q, true)
}