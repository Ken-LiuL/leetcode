package string


type (
	//Stack a stack data structure
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value int
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
func (stack *Stack) Peek() int {
	if stack.length == 0 {
		return 0
	}
	return stack.top.value
}

//Push push new value to the stack
func (stack *Stack) Push(v int) {
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
func (stack *Stack) Pop() int {
	if stack.length == 0 {
		return 0
	}
	top := stack.top
	stack.top = top.prev
	stack.length--
	return top.value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func scoreOfParentheses(S string) int {
	stack := NewStack()
	cur :=  0

	for  i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack.Push(cur)
			cur = 0
		} else {
			cur = stack.Pop() + max(1, cur*2)
		}
	}
	return cur
	
}