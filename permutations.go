package array


func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	res :=  make([][]int, 0, 100)
	for i := 0; i < len(nums); i++ {
		tmp := append([]int{}, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		for _, v := range permute(tmp) {
			 res = append(res, append(v, nums[i]))
		}
	}
    return res
}