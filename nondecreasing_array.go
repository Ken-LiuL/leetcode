package array


func checkPossibility(nums []int) bool {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i+1] < nums[i] {
			if cnt != 0 {
				return false
			}
			if i == 0 || nums[i+1] >= nums[i-1] {
				nums[i] = nums[i+1]
			} else {
				nums[i+1] = nums[i]
			}
			cnt++
		}
	}
	return true
}