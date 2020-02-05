package array

const MaxInt  = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

//there is two group [larger group][smaller group]
func search(nums []int, target int) int {
	lo , hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		num := 0
		//if mid is in smaller group and target is in smaller group
		//or mid is in larger group and target is in larger group
		if (nums[mid] < nums[0]) == (target < nums[0]) {
			 num = nums[mid]
		} else if target < nums[0] {
			//if target is in smaller group, then nums[mid] is in larger group
			num =  MinInt
		} else {
			//if target is in larger group and nums[mid] is in smaller group
			num = MaxInt
		}
		if num < target {
			lo = mid + 1
		} else if num > target{
			hi = mid
		} else {
			return mid
		}
	}
	return -1

}