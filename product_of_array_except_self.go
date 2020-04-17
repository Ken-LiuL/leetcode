package array


//use two array to record before and after
//[a1, a2, a3, ..., an, an+1,, an+2,..., am]
//one = a1*a2*a3...an-1
//two = an*an+1*an+2..am
//multiply each other element-wise
func productExceptSelf(nums []int) []int {
	previous := make([]int, len(nums))
	after := make([]int, len(nums))
	res := make([]int, len(nums))
	previous[0] = 1
	after[len(nums) - 1] = 1
	for i := 1; i < len(nums); i++ {
		 previous[i] =  previous[i-1]*nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		 after[i] = after[i+1]*nums[i+1]
	}

	for i := 0; i < len(nums);i++ {
		res[i] = previous[i]*after[i]
	}
	return res
}