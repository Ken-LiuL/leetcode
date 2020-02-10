package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findMaxConsecutiveOnes(nums []int) int {
	 cnt, m := 0, 0
	 for _, n := range nums  {
		   if n == 1 {
			   cnt++
		   } else  {
			   m = max(cnt, m)
			   cnt = 0
		   }
	 }
	 return max(cnt, m)
}