package hash

func countElements(arr []int) int {
	  cache := make(map[int]int)
	  for _, n := range arr {
		  cache[n] = 1
	  }
	  res := 0
	  for _, n := range arr {
		  if _, ok := cache[n+1]; ok {
				res += 1
		  }
	  }
	  return res
}

