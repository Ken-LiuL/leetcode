package dynamic

func min(a, b, c int) int {
	res := a
	if b < res {
		res = b
	}
	if c < res {
		res = c
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
//https://leetcode.com/problems/maximal-square/discuss/61802/Extremely-Simple-Java-Solution-%3A)
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n, res := len(matrix)+1, len(matrix[0])+1, 0
	dp = make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	//use additional row and col to solve initialization issue
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				res = max(res, dp[i][j])
			}
		}
	}
	return res*res
}