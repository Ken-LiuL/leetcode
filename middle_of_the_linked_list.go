package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	 slow, fast := head, head
	 for fast != nil && fast.Next != nil {
		 fast = fast.Next
		 if fast != nil {
			 fast = fast.Next
		 }
		 slow = slow.Next
	 }
	 return slow
}