package stack

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestValidParentheses(s string) int {
	stack := NewStack()
	maxV, left := 0, -1
	for j := 0; j < len(s); j++ {
		if s[j] == '(' {
			stack.Push(j)
		} else {
			if stack.isEnd() {
				left = j
			} else {
				stack.Pop()
				if stack.isEnd() {
					maxV = max(maxV, j - left)
				} else {
					maxV = max(maxV, j - stack.Peek())
				}
			}
		}
	}
	return maxV
}