package array


func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	var prev []int
	for  i := 0; i < numRows; i++ {
		if i == 0 {
		   prev = []int{1}
		   res = append(res, prev)
		} else if i == 1 {
		   prev = []int{1, 1}
		   res = append(res, prev)
		} else {
		   r := make([]int, i+1)
		   for j := 0; j < i+1; j++ {
			   if j == 0 {
				   r[j] = 1
			   } else if j == i {
				   r[j] = 1
			   } else {
				   r[j] = prev[j]+prev[j-1]
			   }
		   } 
		   res = append(res, r)
			prev = r
		}
	}
	return res
}