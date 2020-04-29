package array

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	start, end := 0, nums[0]
	res := 1
	for end < len(nums)-1 {
		res += 1
		next := 0
		for i := start; i <= end; i++ {
			next = max(i+nums[i], next)
		}
		start, end = end+1, next
	}
	return res
}
