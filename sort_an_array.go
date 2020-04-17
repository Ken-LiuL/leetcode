package array


func findMinimum(nums []int) int {
	mini, miniInd := 100000, 0
	for ind, n := range nums {
		if n < mini {
			mini = n
			miniInd = ind
		}
	}
	return miniInd
}
//selection sort
func sortArray(nums []int) []int {
	arry := make([]int, 0, len(nums))
	cnt := len(nums)
	for cnt > 0 {
		mInd = findMinimum(nums)
		arry = append(arry, nums[mInd])
		nums =  append(nums[:mInd], nums[mInd+1:]...)
		cnt--
	}
	return arry
}

//quick sort
func sortArray2(nums []int) []int {
	  if len(nums) < 2 {
		  return nums
	  }
	  pivot := nums[0]
	  less := make([]int, 0, len(nums))
	  greater :=  make([]int, 0, len(nums))
	  for _, n := range nums[1:] {
		  if n <= pivot  {
			 less = append(less, n)
		  } else {
			  greater = append(greater, n)
		  }
	  }
	  return append(append(sortArray2(less), pivot), sortArray2(greater)...)
}