package array

const MaxInt  = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

//there is two group [larger group][smaller group]
func search(nums []int, target int) int {
	lo , hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		num := 0
		if nums[mid] == target {
			return true
		}
		//the only thing we need to add for this question
		//when nums[mid] == nums[hi - 1]
		//when can have like 1 2 1 1, so in this case, the right part (1 1) is useless
		//when can have like 1 1 1 3 1, so in this case, the right part is not in order
		//because to summarize, we have generally two exceptions:
		//all elements are same for the right smaller part
		//or the right part is out of order
		
		if nums[mid] == nums[hi-1] {
			hi--
			continue
		}
		for mid < hi && nums[mid] 
		//if mid is in smaller group and target is in smaller group
		//or mid is in larger group and target is in larger group
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			num = nums[mid]
		}  else if target < nums[0] {
			//if target is in smaller group, then nums[mid] is in larger group
			num =  MinInt
		} else {
			//if target is in larger group and nums[mid] is in smaller group
			num = MaxInt
		}
		if num < target {
			lo = mid + 1
		} else {
			hi = mid
		} 
	}
	return false

}