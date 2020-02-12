package dynamic

var dp []int = make([]int, 10000)
var INT_MAX = int((^uint(0)) >> 1)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		res  := INT_MAX
		for _, c := range coins {
			if i - c >= 0 { 
				res = min(res, dp[i - c])
			} 
		}
		if res != INT_MAX  {
			dp[i] = res+1
		} else {
			dp[i] = res
		}
	}
	if dp[amount] == INT_MAX {
        return -1
    } else {
        return dp[amount]
    }
	return dp[amount]
}