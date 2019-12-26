package 

type (
	stack  struct {
		top *node
		length int
	}
	node struct {
		next *node
		val *TreeNode
	}
)

func newStack() *stack {
	return &stack{}
}

func (s *stack) push(n *TreeNode) {
	t := &node{val:n, next: s.top}
	s.top = t
	s.length ++
}

func (s *stack) pop() *TreeNode {
	if s.length == 0 {
		return nil
	}
	tmp := s.top
	s.top = tmp.next
	s.length--
	return tmp.val
}


func (s *stack) isEmpty() bool {
	return s.length == 0
}
//use a stack and non-recursive way
func inorderTraversal(root *TreeNode) []int {
	 arr := make([]int, 0, 20)
	 s := newStack()
	 cur := root
	 for cur != nil || !s.isEmpty() {
		 for cur != nil {
			 s.push(cur)
			 cur = cur.Left
		 }
		 cur = s.pop()
		 arr = append(arr, cur.Val)
		 cur = cur.Right
	 }
	 return arr

}
