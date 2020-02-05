package  List


//a lot of linked list problem , you should create a dummy node as first node
//so that to avoid corner case
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	counter := 1
	dummy := &ListNode{Next:head}
	pre := dummy
	//move to the start point
	for counter < m {
		pre = head
		head = head.Next
		counter++
	}
	start = head
	var cur *ListNode 
	//move to the end
	for counter <= n {
		next := head.Next
		head.Next = cur
		cur = head
		head = next
		counter++
	}
	pre.Next = cur
	start.Next = head
	return dummy.Next
}