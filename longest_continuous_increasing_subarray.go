package array

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res, longest := 1, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			res++
		} else {
		   longest = max(longest, res)
			res = 1
		}
	}
	longest = max(longest, res)
	return longest

}