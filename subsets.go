package array


//https://leetcode.com/problems/subsets/discuss/27288/My-solution-using-bit-manipulation
func pow(a, b int) int {
    res := 1
	for i := 0; i < b; i++ {
		 res *= a
	}
	return res
}

func subsets(nums []int) [][]int {
	l := len(nums)
	subsetNum := pow(2, l)
    fmt.Println(l)
	res := make([][]int, subsetNum)
	for k := 0; k < subsetNum; k++ {
		res[k] = make([]int, 0, l)
	} 
	for i := 0; i < l; i++ {
		for j := 0; j < subsetNum; j++ {
 			 if  (j >> i ) & 1 == 1 {
				 res[j] = append(res[j], nums[i])
			 }
		}
	}
	return res
}