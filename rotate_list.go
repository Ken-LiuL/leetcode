package list


func rotateRight(head *ListNode, k int) *ListNode {
	 if head == nil [
		 return head
	 ]
	 len := 0
	 cur := head
	 //get the length of the linked list
	 for cur != nil {
		 cur = cur.Next
		 len++
	 }
	 //get the real movement position
	 k %= len
	 if k == 0 {
		 return head
	 }

	 dummy := &ListNode{Next: head}
	 pre := dummy
	 cur = head
	 for len != k {
		 pre = cur
		 cur = cur.Next
		 len--
	 }
	 dummy.Next = cur
	 pre.Next = nil
	 for cur != nil {
		 pre = cur
		 cur = cur.Next
	 }
	 pre.Next = head
	 return dummy.Next

}