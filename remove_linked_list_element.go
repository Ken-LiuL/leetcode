package list


func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next:head}
	cur := dummy.Next
	prev := dummy
	for  cur != nil {
		if cur.Val == val {
		  prev.Next = cur.Next
        } else {
          prev = cur
        }
        cur = cur.Next
    }
	return dummy.Next
}