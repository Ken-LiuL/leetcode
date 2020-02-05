package sort


func firstMissingPositive(nums []int) int {
	findNumberPositiveNumbers(nums)
	i := 0
	for i != len(nums) {
		if nums[i] != i+1 {
			return i+1
		}
		i++
	}
	return len(nums)+1
}
//use to pointer to move all positive numbers to the front of the array
//this problem is like to put N numbers in  N bins
func findNumberPositiveNumbers(nums []int) {
	i, l := 0, len(nums)
	for i < l {
		//swap nums[ind] and nums[n-1] if n-1 < len(nums)
		if nums[i] > 0 && nums[i] <= l && nums[i] != nums[nums[i]-1]{
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		} else {
			i++
		}
	}
}