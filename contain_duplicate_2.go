package array

func containsNearbyDuplicate(nums []int, k int) bool {
	cache := make(map[int]int)
	for  ind, n := range nums {
		if r, ok := cache[n]; ok {
			 if ind - r  <= k {
				 return true
			 } else {
				 cache[n] = ind
			 }
		} else {
			cache[n] = ind
		}
	}
}