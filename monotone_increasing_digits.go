package string


import "strconv"
func monotoneIncreasingDigits(N int) int {
    arr := []byte(strconv.Itoa(N))
	 ind := len(arr) - 1
	 for i := len(arr) - 1; i > 0; i-- {
		 if arr[i-1] > arr[i] {
			 arr[i-1]--
			 ind = i-1
		 }
	 }
     ind++
	 for ind < len(arr)   {
		 arr[ind] = '9'
		 ind++
	 }
	 trailing := 0 
	 for arr[trailing] == '0' {
		 trailing++
	 }
     res , _ := strconv.Atoi(string(arr[trailing:]))
     return res
}