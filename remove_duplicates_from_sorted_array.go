package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	newLength := 0
	if len(nums) == 0 {
		return newLength
	}
	cur := nums[0]
	newLength++
	for _, n := range nums[1:] {
		if n != cur {
			nums[newLength] = n
			newLength++
			cur = n
		}
	}
	return newLength
}