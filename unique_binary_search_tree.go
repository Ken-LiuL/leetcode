package tree


var dp map[int]int

//dynamic programming 
func numTrees(n int) int {
	dp = make(map[int]int)
	dp[0], dp[1] = 1, 1
	for i :=2 ; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}