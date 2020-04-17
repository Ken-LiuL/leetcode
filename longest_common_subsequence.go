package dynamic

var dp [][]int

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

func longestCommonSubsequence(text1, text2 string) int {
	dp := make([]int, len(text2)+1)
	for i := 1; i < len(text1)+1; i++ {
		next := make([]int, len(text2)+1)
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				next[j] = 1 + dp[j-1]
			} else {
				next[j] = max(max(dp[j], next[j-1]),dp[j-1])
			}
		}
		dp = next
	}
    return  dp[len(text2)]
}
