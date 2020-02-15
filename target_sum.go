package dynamic 


//https://leetcode.com/problems/target-sum/discuss/97334/Java-(15-ms)-C%2B%2B-(3-ms)-O(ns)-iterative-DP-solution-using-subset-sum-with-explanation
func findTargetSumWays2(nums []int, S int) int {
	totalSum := 0
	for _, n := range nums {
		totalSum += n
	}
	if (S + totalSum ) % 2 == 1 {
        return 0
    }
	//out of the possible range 
	if S > totalSum || S < -totalSum {
		return 0
	}
	return subsetSum(nums, (S + totalSum)/2)
}

//find number of combinations to get the sum of target
func subsetSum(nums []int, target int) int {
	 dp := make([]int, target+1)
	 dp[0] = 1
	 for _, n := range nums {
		 for i := target; i >= n; i-- {
			 dp[i] += dp[i - n]
		 }
	 }
	 return dp[target]
}


//https://leetcode.com/problems/target-sum/discuss/97335/Short-Java-DP-Solution-with-Explanation
func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if S > sum || S < -sum {
		return 0
	}
	dp := make([]int, 2*sum+1)
	dp[0+sum] = 1

	for i := 0; i < len(nums); i++ {
		next := make([]int, 2*sum + 1)
		for j := 0; j < 2*sum+1; j++ {
			if j+nums[i] < 2*sum + 1 {
				next[j] += dp[j+nums[i]] 
			}
			if j-nums[i] >= 0 {
				next[j] += dp[j-nums[i]]
			}
		}
		dp = next
	}
	return dp[sum + S]
}