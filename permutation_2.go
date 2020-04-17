package array
import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return helper(nums)
}

func helper(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	res :=  make([][]int, 0, 100)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		tmp := append([]int{}, nums[:i]...)
		tmp = append(tmp, nums[i+1:]...)
		for _, v := range helper(tmp) {
			 res = append(res, append(v, nums[i]))
		}
	}
    return res
}