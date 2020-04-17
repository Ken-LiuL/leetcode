package string

type (
	//Stack a stack data structure
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value rune
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
func (stack *Stack) Peek() rune {
	if stack.length == 0 {
		return rune(' ')
	}
	return stack.top.value
}

//Push push new value to the stack
func (stack *Stack) Push(v rune) {
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
func (stack *Stack) Pop() rune  {
	if stack.length == 0 {
		return rune(' ')
	}
	top := stack.top
	stack.top = top.prev
	stack.length--
	return top.value
}

func removeDuplicateLetters(s string) string {
	stack := NewStack()
	seen := make([]int, 26)
	last := make([]int, 26)
	for i, v := range s {
		last[v - 'a'] = i
	}

	for i, v := range s {
		if seen[v - 'a'] != 0 {
			continue
		}
		for !stack.isEnd() && (i < last[stack.Peek() - 'a'] && v < stack.Peek()) {
			seen[stack.Pop() - 'a'] = 0
		}
		seen[v - 'a']++
		stack.Push(v)
	}
	res := make([]rune, 0, 100)
	for !stack.isEnd() {
		res = append(res, stack.Pop())
	}
	l := len(res)
	for i := 0; i <  l / 2; i++ {
		 res[i], res [l-i-1] = res[l-i-1], res[i]
	}
	return string(res)
}