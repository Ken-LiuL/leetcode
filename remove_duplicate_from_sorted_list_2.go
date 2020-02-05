package list



func deleteDuplicates(head *ListNode) *ListNode {
   dummy := &ListNode{Next: head}
   cur := dummy.Next
   pre := dummy
   flag := false

   for cur != nil {
	   if cur.Next != nil && cur.Next.Val == cur.Val {
		  cur.Next = cur.Next.Next
		  flag  = true
	   } else { 
		  if flag {
			  pre.Next = cur.Next
              cur = cur.Next
			  flag = false
          } else {
		    pre = cur
	   	    cur = cur.Next
	      }
        }
}
       return dummy.Next

}