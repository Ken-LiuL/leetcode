package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestOnes(A []int, K int) int {
	 start, end := 0, 0
	 zind  := make([]int, 0, len(A))
	 longest := 0
	 for start < len(A) && end < len(A) {
		 if A[end] == 0 {
			 zind = append(zind, end)
			if len(zind) > K  {
			  longest = max(longest, end - start)
			  start = zind[0]+1
			  zind = zind[1:]
			} 
		 }
		 end++
	 }
	 return max(longest, end - start)
}