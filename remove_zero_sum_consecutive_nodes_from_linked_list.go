package list


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeZeroSumSublists(head *ListNode) *ListNode {
	cache := make(map[int]*ListNode)
	dummy := &ListNode{Next: head, Val: 0} 
	h ,  sum := dummy, 0
	for h != nil {
		sum += h.Val
		if _, ok := cache[sum]; ok {
		   cur := cache[sum].Next
		   p := cur.Val+sum	 
		   for p != sum {
			   delete(cache, p)	
			   cur = cur.Next
			   p += cur.Val
		   }
		   cache[sum].Next = cur.Next
		} else {
		  cache[sum] = h
		}
		h = h.Next
	}
	return dummy.Next
}