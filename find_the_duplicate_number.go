package array

func findDuplicate(nums []int) int {
	fast, slow := nums[nums[0]], nums[0]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	find := 0

	fast = 0
	for fast != slow {
		fast = nums[fast]
		slow = nums[slow]
	}
	return slow

}