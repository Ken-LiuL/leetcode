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

func smallestSubsequence(text string) string {
	stack := NewStack()
	last := make([]int, 26)
	for ind, t := range text {
		last[t-'a'] = ind
	}
	seen := make([]int, 26)

	for i, t := range text {
		if seen[t - 'a']  != 0 {
			seen[t - 'a']++
			continue
		}
        for !stack.isEnd() && (t < stack.Peek() && i < last[stack.Peek()-'a'] ) {
			seen[stack.Pop()-'a'] = 0
		}
        		seen[t-'a']++
		stack.Push(t)
	} 
	res := make([]rune, 0, 100)
	for !stack.isEnd() {
		res = append(res, stack.Pop())
	}
	l := len(res)
	for i := 0; i < l/2; i++  {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return string(res)
}