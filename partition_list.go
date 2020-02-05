package list

func partition(head *ListNode, x int) *ListNode {
	dummy := &ListNode{Next: head}
	p := dummy
	tail := dummy
	//move p to the first node larger than x
	for p != nil && p.Next != nil && p.Next.Val < x {
	   tail = tail.Next
	   p = p.Next
   }
	for p != nil && p.Next != nil {
		if p.Next.Val >= x {
			p = p.Next
		} else {
			tmp := p.Next
			p.Next = tmp.Next
			tmp.Next = tail.Next
			tail.Next = tmp
			tail = tmp
		}
	}
	return dummy.Next
}