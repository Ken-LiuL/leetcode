package dynamic


var memorization [][]int 

func uniquePaths(m int, n int) int {
	memorization = make([][]int, n)
	for i := 0; i < n; i++ {
		memorization[i] = make([]int, m)
	}
	return helper(m, n)
}

func helper(m int, n int) int {
	if m == 1 || n == 1{
		return 1
	}
    if memorization[n-1][m-1] == 0 {
	    memorization[n-1][m-1] =  helper(m - 1, n)+helper(m, n - 1)
    }
    return memorization[n-1][m-1]
}