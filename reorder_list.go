package list

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
	}
	return cur
}
//split the linked list as two part [part1][part2]
//reverse the second half [part1][reversed-part2]
//traverse these two parts, and insert each of the second part into the first one
func reorderList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//split by using fast and low pointers
	fast, slow := head, head
	steps_count, steps_limit := 0, 2
	for fast != nil {
		fast = fast.Next
		steps_count++
		if steps_count == steps_limit {
			slow = slow.Next
			steps_count = 0
		}
	}
	//start to reverse list from slow to fast pointer
	reversedHead := slow.Next
	slow.Next = nil
	reversedHead = reverseList(reversedHead)

	left, right := head, reversedHead 
	//length of left part will be larger or equal to the right part
	for right != nil {
		leftNext := left.Next
		left.Next = right
		left = leftNext
		rightNext := right.Next
		right.Next = leftNext
		right = rightNext
	}
	return head
}