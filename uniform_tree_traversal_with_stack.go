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

func (s *stack) seek() *TreeNode {
	if s.length == 0 {
		return nil
	}
	return s.top.val
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

func traversal(root *TreeNode) []int {
	arr := make([]int, 0, 10)
	cur := root
	s := newStack()
	pre := cur
	for cur != nil || !s.isEmpty() {
		for cur != nil {
			//set as preorder
			//arr = append(arr, cur.Val) 
			s.push(cur)
			cur = cur.Left
		}                               
		cur  = s.seek()  
		//set as inorder
		//if cur.Right != pre {
		//	arr = append(arr, cur.Val)
	    //}
		cur = cur.Right 
		if cur == nil || cur == pre {
			pre = s.pop()   
			//postorder
			//arr = append(arr, pre.Val)
			cur = nil
		} 	
	}
	return arr
}