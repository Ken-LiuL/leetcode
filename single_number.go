package array

// a ^ a = 0
// 0 ^ anything = anything
// 1 ^ anything = ~anything
func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^=  nums[i]
	}
	return nums[0]
}