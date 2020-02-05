package list


//use two pointer to move 
func swapPairs(head *ListNode) *ListNode {
	// when length of the list is <= 1
	 if head == nil || head.Next == nil {
		 return head
	 }
	 cur := head.Next
	 pre := head
	 start := cur
	 for  cur != nil {
          next, nextNext := cur.Next, cur.Next
          if next != nil && next.Next != nil {
             nextNext = cur.Next.Next
          }
		  cur.Next = pre
		  pre.Next =  nextNext
		  cur = nextNext
		  pre = next
     }
	 return start
}