package list

// it is the same as add two number, except, in this case the digit is normal order
// like for 1234 , the linked list is 1--> 2 --> 3-->4

type (
	//Stack a stack data structure
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

//NewStack return new stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

//Len stack length
func (stack *Stack) Len() int {
	return stack.length
}

//Peek return top value of this stack
func (stack *Stack) Peek() interface{} {
	if stack.length == 0 {
		return nil
	}
	return stack.top.value
}

//Push push new value to the stack
func (stack *Stack) Push(v interface{}) {
	newTop := &node{v, stack.top}
	stack.top = newTop
	stack.length++
}

//isEnd check whether it is the end of stack
func (stack *Stack) isEnd() bool {
	if stack.length == 0 {
		return true
	} else {
		return false
	}
}

//Pop pop the current top value
func (stack *Stack) Pop() interface{} {
	if stack.length == 0 {
		return nil
	}
	top := stack.top
	stack.top = top.prev
	stack.length--
	return top.value
}

//ListNode the node of linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := NewStack()
	stack2 := NewStack()

	notEnd := true

	for notEnd {
		if l1 == nil && l2 == nil {
			notEnd = false
		}
		if l1 != nil {
			stack1.Push(l1.Val)
			l1 = l1.Next
		}
		if l2 != nil {
			stack2.Push(l2.Val)
			l2 = l2.Next
		}
	}

	inc, res := addTwoNumbersHelper(stack1, stack2)
	if inc != 0 {
		return &ListNode{inc, res}
	}
	return res

}

func addTwoNumbersHelper(stack1 *Stack, stack2 *Stack) (int, *ListNode) {
	inc := 0
	var next *ListNode
	var rest *Stack
	for !stack1.isEnd() && !stack2.isEnd() {
		sum := stack1.Pop().(int) + stack2.Pop().(int) + inc
		next = &ListNode{sum % 10, next}
		if sum >= 10 {
			inc = 1
		} else {
			inc = 0
		}
	}
	if !stack1.isEnd() {
		rest = stack1
	} else {
		rest = stack2
	}
	for !rest.isEnd() {
		sum := rest.Pop().(int) + inc
		next = &ListNode{sum / 10, next}
		if sum >= 10 {
			inc = 1
		} else {
			inc = 0
		}
	}
	return inc, next
}
