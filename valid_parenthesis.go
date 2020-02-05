package stack


func isValid(s string) bool {
	stack := NewStack()   
	for _, n := range s {
		if n == '(' || n == '[' || n == '{' {
			stack.Push(n)
		} else {
			if stack.Peek() != nil {
			   v :=  stack.Pop()
			   if (v == '(' && n ==  ')') || (v == '[' && n == ']') || (v == '{' && n == '}') {
				   continue
			   } else {
				   return false
			   }
			} else {
				return false
			}
		}
	}
	 if stack.isEnd() {
		 return true
	 } else {
		 return false
	 }
	 
 }