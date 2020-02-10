package array

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func findShortestSubArray(nums []int) int {
	 cnt := make(map[int]int)
	 first := make(map[int]int)
	 res, degree := 1, 1
	 for i := 0; i < len(nums); i++ {
		 if _, ok := cnt[nums[i]]; ok {
			  cnt[nums[i]]++
			  if degree < cnt[nums[i]] {
				  degree = cnt[nums[i]]
				  res = i - first[nums[i]] + 1
			  } else if degree == cnt[nums[i]]{
				  res =  min(res, i - first[nums[i]]+1)
			  }
		 } else {
			 cnt[nums[i]] = 1
			 first[nums[i]] = i
		 }
	 }
	 return  res
}