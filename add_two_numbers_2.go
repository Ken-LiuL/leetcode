package list

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := NewStack()
	stack2 := NewStack()
	for l1 != nil {
		stack1.Push(l1)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2.Push(l2)
		l2 = l2.Next
	}

	carries := 0
	dummy := &ListNode{}
	for !stack1.isEnd() && !stack2.isEnd() {
		  n1, n2 := stack1.Pop().(*ListNode), stack2.Pop().(*ListNode)
		  sum := n1.Val + n2.Val + carries
		  carries = sum / 10
		  node := &ListNode{Next: dummy.Next, Val: sum % 10}
		  dummy.Next = node
	}
	s := stack1
	if s.isEnd() {
		s = stack2
	}
	for !s.isEnd() {
		n1 := s.Pop().(*ListNode)
		sum := n1.Val + carries
		carries = sum / 10
		node := &ListNode{Next: dummy.Next, Val: sum % 10} 
		dummy.Next = node
	}
	if carries != 0 {
        node := &ListNode{Next: dummy.Next, Val: 1}
        dummy.Next = node
    }
	return dummy.Next
}