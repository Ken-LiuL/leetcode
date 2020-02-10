package array


func findPairs(nums []int, k int) int {
	if k < 0 {
	  return 0 
	}
	res := 0 
	cache := make(map[int]int)
	for _, n := range nums {
		//jump duplicate element
		if cnt, ok := cache[n]; ok {
			if k == 0 && cnt == 1 {
				res++
			} 
			cache[n]++
		} else {
			if _, ok := cache[n - k]; ok {
				res++
			}
			if _, ok := cache[n + k]; ok {
				res++
			}
			cache[n] = 1
		} 
	}
	return res
}