package array

func subarraySum(nums []int, k int) int {
	dict := make(map[int]int)
	dict[0] = 1
	res, sum := 0, 0
	for _, n := range nums {
		sum += n
		if v, ok := dict[sum-k]; ok {
			res += v
		}
		if _, ok := dict[sum]; ok {
			dict[sum] += 1
		} else {
			dict[sum] = 1
		}
	}
	return res
}
