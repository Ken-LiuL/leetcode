package array

func findDisappearedNumbers(nums []int) []int {
	//place it in the right place
	for i := 0; i < len(nums); {
		if  nums[i] != i+1 && nums[nums[i] - 1] != nums[i] {
			nums[i],  nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		} else {
			i++
		}
	}
	res := make([]int, 0, 3)
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}