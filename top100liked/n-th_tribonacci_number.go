package recursion

//0 <= n <= 37
//for record value of tribonacci(n)
var memorization = [38]int{0, 1, 1}

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if memorization[n] != 0 {
		return memorization[n]
	}
	memorization[n] = tribonacci(n-3) + tribonacci(n-2) + tribonacci(n-1)
	return memorization[n]
}

