package dynamic

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//https://leetcode.com/problems/house-robber/discuss/156523/From-good-to-great.-How-to-approach-most-of-DP-problems.
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
	}
	now, prev := 0, 0
	
	for i := len(nums) - 1; i >= 0; i-- {
		now, prev = max(nums[i] + prev, now), now
	}
	return now
}
