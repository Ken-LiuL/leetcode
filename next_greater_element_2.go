package  list

func nextGreaterElements(nums []int) []int {
	 stack := make([]int, 0, len(nums) * 2)
	 res := make([]int, len(nums))
	 
	 for i, n := range nums {
           for len(stack) != 0 &&  nums[stack[len(stack)- 1]] < n {
			    res[stack[len(stack) - 1]] = nums[i]
				stack = stack[: len(stack) - 1]
		   } 
		   stack = append(stack, i)
	 }
	 for i, n := range nums {
		   for len(stack) != 0 && nums[stack[len(stack) - 1]] < n {
				res[stack[len(stack)-1]] = nums[i]
				stack = stack[:len(stack) - 1]
		   }
	 }
	 for _, v := range stack {
		 res[v] = -1
	 }
	 return res
}