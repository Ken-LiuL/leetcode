package hash

func max(a, b int) int {
	if a > b {
		return a 
	}
	return b
}

//https://leetcode.com/problems/contiguous-array/discuss/99655/Python-O(n)-Solution-with-Visual-Explanation
func findMaxLength(nums []int) int {
	 d := make(map[int]int)
	 m := 0
	 sum := 0
     d[0] = -1
	 for idx, num := range nums {
		 if num == 0 {
			sum -= 1
		 } else {
			 sum += 1
		 }
		 if previous, ok := d[sum]; ok {
			m = max(m, idx - previous)
		 } else {
			d[sum] = idx
		 }
	 }
	 return m
}
