package dynamic


var memorization []int = make([]int, 10000)
const INT_MAX = int(^uint(0) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func numSquares(n int) int {
	 memorization[0] = 0
	 memorization[1] = 1

	 for i := 2; i <= n; i++ {
		res := INT_MAX
		for j := 1; j*j <= i; j++ {
			res = min(res, memorization[i - j*j])
		}	 
		memorization[i] = res + 1
	 }
	 return memorization[n]
}