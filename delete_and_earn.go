package dynamic


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func deleteAndEarn(nums []int) int {
	  values := make([]int, 10000)
	  for _, n := range nums {
		  values[n-1] += n
	  }
	  now, pre := 0, 0
	  for i := len(values) - 1; i >= 0; i-- {
		  now, prev := max(prev+values[i], now), now
	  }
	  return now
}