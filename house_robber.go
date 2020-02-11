package dynamic

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

var memorization map[int]int

func rob(nums []int) int {
	memorization = make(map[int]int)
	return helper(nums, 0)
}
func helper(nums []int, start int) int {
	if len(nums) == 0 {
        return 0
    }
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if v, ok := memorization[start]; ok {
		return v
	} else {
		memorization[start] = max(nums[0] + helper(nums[2:], start+2), nums[1] + helper(nums[3:], start+3))
		return memorization[start]
	}
}

