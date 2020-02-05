package list



//use two pointer to detect length of the loop first
//then we can know the start of the cycle
//use http://www.siafoo.net/algorithm/11
// the logic is like this:
// we bring the slow pointer  to the loop as soon as possbile
// then we use faster pointer to iterate, when the iteration length >= loop length
// then it will eventually met the lower point, and 
// at the same time, we could know the length of the loop
func detectCycle(head *ListNode) *ListNode {
	if head == nil  || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	steps_taken, step_limit := 0, 2 
	for {
        if fast == nil {
            return nil
        }
		fast  = fast.Next
		steps_taken++ 
		if fast == slow {
			break
		} 
		if steps_taken == step_limit {
			steps_taken = 0
			step_limit *= 2
			slow = fast
		}
	 }  
	 fast, slow = head, head
	for steps_taken > 0 {
		fast = fast.Next
        steps_taken--
	}
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow 
}
 



