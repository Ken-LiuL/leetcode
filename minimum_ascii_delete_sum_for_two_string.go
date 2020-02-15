package dynamic

func min(a, b, c int) int {
	m := a
	if b < a {
		m  = b
	}
	if c < m {
		m = c
	}
	return m
}

func minimumDeleteSum(s1 string, s2 string) int {
	m,  n := len(s1)+1, len(s2) + 1
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n ; i++ {
		dp[i] = dp[i-1] + int(s2[i-1])
	}
	for i := 1; i < m ;i++ {
		next := make([]int, n)
		next[0] = dp[0] + s1[i-1]
		for j:= 1; j < n; j++ {
			if s1[i-1] == s2[j-1] {
				next[j] = dp[j-1]
			} else {
				c1, c2 := int(s1[i-1]), int(s2[j-1])
				next[j] = min(dp[j]+c1, next[j-1]+c2, dp[j-1]+c1+c2)
			}
		}
		dp = next
	}
	return dp[n-1]
}