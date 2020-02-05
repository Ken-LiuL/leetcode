package tree

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

func preorderTraversal(root *TreeNode) []int {
	arr := make([]int, 0, 10)
	cur := root
	s := newStack()

	for cur != nil || !s.isEmpty() {
		for cur != nil {
			s.push(cur)
			arr = append(arr, cur.Val)
			cur = cur.Left
		}
		cur = s.pop()
		cur = cur.Right
	}

}

