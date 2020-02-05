package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	newLength := 0
	if len(nums) <= 1 {
		return newLength
	}
	pre := nums[0]
	cur := nums[1]
	newLength = 2
	for _, n := range nums[2:] {
		if n != cur || n != pre {
			nums[newLength] = n
			newLength++
			pre, cur = cur, n
		} 
	}
	return newLength
}