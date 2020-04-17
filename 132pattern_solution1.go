package list

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//https://medium.com/brain-framework/132-pattern-1b890c763bd5
func find132pattern(nums []int) bool {
    stack := make([][]int,0, len(nums))

	 for _, n := range nums { 
		 if len(stack) == 0 || n < stack[len(stack) - 1][0]{
			stack = append(stack, []int{n, n})
		 } else{
             mini := stack[len(stack) - 1][0]
			for len(stack) != 0 && stack[len(stack) - 1][1] < n {
				stack = stack[:len(stack) - 1]
			}
			if len(stack) != 0 && stack[len(stack) - 1][0] < n && stack[len(stack) - 1][1] > n {
				return true
			}
			stack = append(stack, []int{mini, n})
		 }
	 }
	 return false
}