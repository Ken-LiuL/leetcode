package array



func intersect(nums1 []int, nums2 []int) []int {
	cache := make(map[int]int)
	res := make([]int, 0, 10)

	for _, n := range nums1 {
		cache[n]++
	}
	for _, n := range nums2 {
		if v, ok := cache[n]; ok && v > 0 {
			res = append(res, n)
			cache[n]--
		}
	}
	return res
}