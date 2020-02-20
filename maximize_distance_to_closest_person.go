package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func maxDistToClosest(seats []int) int {
	l := len(seats) 
	start := -1
	res := 1
	for ind, s := range seats {
		 if s == 1 {
			 if start == -1 {
				 start = ind
				 res = max(res, start)
			 } else {
				 res = max(res,  (ind - start) / 2)
				 start = ind
			 }
		 }
	}
	if start != l - 1 {
		res = max(res, (l-1-start))
	}
    return res
}