package list


func deleteDuplicates(head *ListNode) *ListNode {
      cur := head
	 for cur != nil {
		 if cur.Next != nil && cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		 } 
		 cur = cur.Next
	 }
	 return head
}