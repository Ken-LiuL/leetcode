package recursion

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	//if sum of numbers in the array can not be divided by k, then , it should return false
	if sum%k != 0 {
		return false
	}
	//the target of subset sum
	target := sum / k

	sort.Ints(nums)

	subset := make([]int, 0, k)

	return canPartitionHelper(nums, k, len(nums)-1, subset, target)

}

func canPartitionHelper(nums []int, k int, start int, subset []int, target int) {
	if start < 0 {
		return true
	}
	for i := 0; i < k; i++ {
		//case 1:  nums[start] > subset[i-1]
		//case 2: put a number in subset[i-1] could not work, so [..., nums[start], 0, 0, 0just would not work
		//   then [...,0, nums[start]] would met the same problem
		if i > 0 && subset[i-1] == 0 {
			return false
		}
		if subset[i]+nums[start] <= target {
			//put start-th element to the i-th drawer
			subset[i] += nums[start]
			//check rest part
			if canPartitionHelper(nums, k, start-1, subset, target) {
				return true
			}
			//not fit
			//take start-th element out of i-th drawer
			subset[i] -= nums[start]
		}
	}
	return false
}
