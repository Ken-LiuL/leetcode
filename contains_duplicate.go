package array

import "sort"

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
        return false
    }
	sort.Ints(nums)
	prev := nums[0]
	for _, n := range nums[1:] {
		if prev == n {
			return true
		} else {
			prev = n
		}
	}
	return false
}