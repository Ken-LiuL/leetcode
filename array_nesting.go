package array

func arrayNesting(nums []int) int {
	longest,  limit := 0, 20000 	
	for i := 0; i < len(nums); i++ {
		ind, cnt := i, 0
		for  nums[ind] <  limit {
			cnt++
			nums[ind] += limit
			ind = nums[ind] - limit
		}
		if cnt > longest {
			longest = cnt
		}
	}
	return longest
}

