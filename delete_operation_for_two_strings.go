package string

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance(word1 string, word2 string) int {
	dp := make([]int, len(word2)+1)
    for j := 1; j < len(word2)+1; j++ {
		dp[j] = j
	}
	for i := 1; i < len(word1) + 1; i ++ { 
		next := make([]int, len(word2)+1) 
        next[0] = i
		for j := 1; j < len(word2)+1; j++ {
			if word1[i-1] ==  word2[j-1] {
				next[j] = dp[j-1]
			} else {
				next[j] = 1+min(dp[j], next[j-1])
            }
		}
		dp = next
	}
	return dp[len(word2)]
}