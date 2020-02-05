package sort


//binary parition + sort two lists
func mergeKLists(lists []*ListNode) *ListNode {
	return partition(lists, 0, len(lists)-1)
}

func partition(lists []*ListNode, start, end int) ListNode {
	if start == end {
		return lists[start]
	}
	if start < end {
		q := start + (end - start)/2
		l1 :=  partition(lists, start, q)
		//q must add 1 to avoid infinit recursion
		l2 :=  partition(lists, q+1, end)
		return mergeTwoLists(l1, l2)
	} else {
		return nil
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var mergedHead *ListNode
	if l1.Val < l2.Val {
		mergedHead = l1
		mergedHead.Next = mergeTwoLists(l1.Next, l2)
	} else {
		mergedHead = l2
		mergedHead.Next = mergeTwoLists(l1, l2.Next)
	}

	return mergedHead
}
