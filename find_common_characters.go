package array

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func commonChars(A []string) []string {
	dp := make([][]int, len(A))
	strs := make([]string, 0, 26)

	for i := 0; i < len(A); i++ {
		dp[i] = make([]int, 26)
		for j := 0; j < len(A[i]);j++ {
			dp[i][A[i][j]-'a']++
		}
	}
	for i := 0; i < 26; i++ {
		res := true
		v := 101
		for j := 0; j < len(A); j++ {
			v = min(v, dp[j][i])
			if v == 0 {
				res = false
				break
			}
		}
		if res {
			for v > 0 {
				v--
				strs = append(strs, string(i+'a'))
			}
		}
	}
	return strs
}