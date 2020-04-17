package list

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0, len(T))
	res := make([]int, len(T))
	for i , v := range T {
	   for len(stack) != 0 && v > T[stack[len(stack) - 1]] {
		   res[stack[len(stack) - 1]] = i - stack[len(stack)-1]
		   stack = stack[:len(stack) - 1]
	   }
	   stack = append(stack, i)
	}
	return res
}