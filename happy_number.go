package number

func sumOfDigits(n int) int {
	s := 0
	for  n > 0 {
		 rem := n % 10
		 s += rem * rem
		 n = n / 10
	}
	return s
}

func isHappy(n int) bool {
	visited := make(map[int]int)
	for {
		if _, ok := visited[n]; ok {
			return false
		}
		visited[n] = 1
		n = sumOfDigits(n)
		if n == 1 {
			return true
		}
	}
}