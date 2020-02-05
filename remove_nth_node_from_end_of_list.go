package list


//use two pointer to traverse the list
//if the difference of the two pointer is n
//then when one pointer move to the end of the list
//then the otehr pointer point exactly to the target node
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	  dummy := &ListNode{Next: head}
	  pre := dummy
	  cur := dummy.Next
	  
	  //two pointer
	  fast :=  cur
	  slow :=  cur

	  //move fast first
	  for n > 0 {
		  fast = fast.Next
		  n--
	  }
	  //move them simutaneously
	  for fast != nil {
		  pre = slow
		  slow = slow.Next
		  fast = fast.Next
	  }

	  //now fast reach the end of the list
	  //slow point to the target node

	  pre.Next = slow.Next
	  return dummy.Next

}
