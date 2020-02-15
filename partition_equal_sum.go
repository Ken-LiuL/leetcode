package dynamic


//https://leetcode.com/problems/partition-equal-subset-sum/discuss/90592/01-knapsack-detailed-explanation
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum % 2 == 1 {
		return false
	}
	sum /= 2
	dp :=  make([]bool, sum+1)
	dp[0] = true
	for _, n := range nums {
		for j := sum; j >= n; j-- {
			dp[j] = dp[j] || dp[j - n]
		}
	}
	return dp[sum]
}