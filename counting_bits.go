package dynamic 


var dp []int 

func countBits(num int) []int {
	dp := make([]int, num + 1)
	dp[0] = 0
	//for example 11011
	//dp(11011) = dp(1101) + (1 % 2)
	for i := 1; i <= num; i++ {
		dp[i] = dp[i >> 1] + i & 1
	}
	return dp
}