package list


//two pointer
//fast runner and slow runner
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}