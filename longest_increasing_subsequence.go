package dynamic


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	longest := 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if(nums[i] > nums[j]) {
				dp[i] = max(dp[i], dp[j]+1)
				longest =  max(dp[i], longest)
			}
		}
	}
	return longest

}