package main

//ListNode the node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)
}

/**
*  iteratively add digit in each node and pass whether the sum overflow 10 to the next iteration
*
*/
func helper(l1 *ListNode, l2 *ListNode, inc int) *ListNode {
	if l1 == nil && l2 == nil {
		if inc == 0 {
			return nil
		}
		return &ListNode{inc, nil}
	}
	l1Val := 0
	l2Val := 0
	var l1Next *ListNode
	var l2Next *ListNode

	if l1 != nil {
		l1Val = l1.Val
		l1Next = l1.Next
	}
	if l2 != nil {
		l2Val = l2.Val
		l2Next = l2.Next
	}
	sum := l1Val + l2Val + inc
	if sum >= 10 {
		inc = 1
	} else {
		inc = 0
	}
	return &ListNode{sum % 10, helper(l1Next, l2Next, inc)}
}
