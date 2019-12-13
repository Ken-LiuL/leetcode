package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *  a typical recursion problem
 * }
 */
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
