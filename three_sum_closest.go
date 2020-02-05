package array
import "sort"

func abs(a int ) {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	//sort first
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums) - 2; i++ {
		second, third, remain := i+1, len(nums)-1, target - nums[i]
		//jump consecutive numbers
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for second < third {
			sum := nums[second] + nums[third] 
			if sum == remain {
				return  target
			}
			if abs(remain - sum) < abs(target - closest) {
				closest =  sum + nums[i]
			}
			if sum > remain {
				third--
			} else {
				second++
			}
		}
	}
	return closest
}

