package dynamic

var memorization map[int]int

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	memorization = make(map[int]int)
	memorization[1] = 1
	memorization[2] = 2
	return helper(n)
}

func helper(n int) int {
	if v, ok := memorization[n]; ok {
		return v
	} else {
		memorization[n] = helper(n-1)+helper(n-2)
		return memorization[n]
	}
}