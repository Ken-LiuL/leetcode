package list


const MOD = 1000000007

func sumSubarrayMins(A []int) int {
	stack := make([]int, 0, len(A))
    previousLess := make([]int, len(A))
    nextLess := make([]int, len(A))

	for i , v := range A {
		for len(stack) != 0  && A[stack[len(stack) - 1]] > v {
			stack = stack[: len(stack) - 1]
		}
		if len(stack) == 0 {
			previousLess[i] = i+1
		} else {
			previousLess[i] = i-stack[len(stack) - 1]
		}
		stack = append(stack, i)
	}
	stack = make([]int,0, len(A))
	for i , v := range A {
		for len(stack) != 0 && A[stack[len(stack) - 1]] > v {
			 nextLess[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			 stack = stack[:len(stack) - 1]
		}
		stack = append(stack, i)
	}
	for _, s := range stack {
		nextLess[s] = len(A) - s
	}

	res := 0
	for i := 0; i < len(A); i++ {
		p, n := previousLess[i], nextLess[i]
		res += p*n*A[i]
	}
	return res % MOD
}