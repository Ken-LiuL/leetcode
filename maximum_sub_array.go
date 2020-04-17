package array

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//Kadane's algorithm
//https://thirumal.blog/2018/03/18/kadane-deep-dive/
func maxSubArray(nums []int) int {
	  maxLocal, maxGlobal := nums[0], nums[0]
	  for _, n := range nums[1:] {
		  maxLocal = max(maxLocal + n, n)
		  maxGlobal = max(maxGlobal, maxLocal)
	  }
	  return maxGlobal
}

