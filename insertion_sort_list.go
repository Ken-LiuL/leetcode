package sort

/*
 * 1.Insertion sort iterates, consuming one input element each repetition, and growing a sorted output list.
 * 2. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list, and inserts it there.
 * 3. It repeats until no input elements remain.
**/
//the key point is to use a pre pointer as help
func insertionSortList(head *ListNode) *ListNode {
	//helper node 
	start := &ListNode{}
	cur := head
	pre :=  start
	for cur != nil {
		//only start from the beginning when current node value is smaller than pre 
		//otherwise, if cur.Val < pre.Next.Val, then each node before pre should be smaller
		//than cur, so there is no need for moving pre to the front
		if pre.Next == nil || pre.Next.Val > cur.Val {
			pre = start
		}
		for pre.Next != nil &&  pre.Next.Val < cur.Val {
			pre = pre.Next
		}
		//preserve next
		next := cur.Next
		//pre.Next is either null or the next node larger than cur.Val
		pre.Next, cur.Next = cur, pre.Next
		//set cur as next
		cur = next
	}

	return start.Next
}