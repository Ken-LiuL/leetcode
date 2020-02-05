package list

func reverseList(head *ListNode) *ListNode {
	cur := &ListNode{}
	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}
	return cur->next
}

func reverseHelper(head *ListNode, newHead *ListNode) *ListNode {
	if head == nil {
		return newHead
	}
	next := head.Next
	head.Next = newHead
	return reverseHelper(next, head)
}

func reverseList2(head *ListNode) *ListNode {
	return reverseHelper(head, nil)
}