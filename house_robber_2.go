package dynamic

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


//https://leetcode.com/problems/house-robber/discuss/156523/From-good-to-great.-How-to-approach-most-of-DP-problems.
func rob1(nums []int) int {
	now, prev := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		now, prev = max(nums[i] + prev, now), now
	}
	return now
}


func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
    if len(nums) == 1 {
        return nums[0]
    }
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(nums[0]+rob1(nums[2:len(nums)-1]), rob1(nums[1:]))
}
