package array



func intersection(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	res := make([]int, 0, 10)

	for _, n := range nums1 {
		cache[n] = 1
	}
	for _, n := range nums2 {
		if _, ok := cache[n]; ok {
			res = append(res, n)
			delete(cache, n)
		}
	}
	return res
}