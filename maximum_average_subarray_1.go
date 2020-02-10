package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxAverage(nums []int, k int) float64 {
	start, end := 0, k-1
	res, maxV := 0, -10000000
	for i := 0; i < k; i++ {
		res += nums[i]
	}
	maxV = max(res, maxV)
	 
	for end < len(nums)  {
		end++
		if end > len(nums) - 1 {
			break
		}
		res += nums[end]
		res -= nums[start]
		start++
		maxV = max(res, maxV)
	} 
   return float64(maxV) / float64(k)

}