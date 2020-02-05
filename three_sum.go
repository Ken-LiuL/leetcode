package array
import "sort"

func threeSum(nums []int) [][]int {
	//sort first
	sort.Ints(nums)
	res := make([][]int, 0, 10)
	//since after sort, the number will from smaller to larger
	for i := 0; i < len(nums) -2 ; i ++ {
		lo, hi, sum := i+1, len(nums)-1, 0- nums[i]
	
		//since sorted, if nums[i] >0, then lo[i] >0 and hi[i] > 0
		//so overall sum will be larger than 0
		if nums[i] > 0 {
			return res
		}
		for lo < hi {
			if(nums[lo]+nums[hi] == sum) {
				res = append(res, []int{nums[i], nums[lo], nums[hi]})
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}
				lo++
				hi--
			} else if(nums[lo]+nums[hi] < sum){
				lo++
			} else {
				hi--
			}
		}
		//if consecutive just continue
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
