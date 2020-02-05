package str

func insertSpecial(s string) []rune {
	//insert # firstly to resolve odd and even issue
	arr := make([]rune, 0, len(s))
	for _, n := range s {
		arr = append(arr, '#')
		arr = append(arr, n)
	}
	arr = append(arr, '#')
	return arr
}
//we use manacher algorithm
func longestPalindrome(s string) string {
	//insert special '#'
	arr := insertSpecial(s)
	//middle out 
	max, maxInd := 0, 0
	for ind, _ := range arr {
		start, end, count := ind-1, ind+1, 0
		for start >=0 && end < len(arr) && arr[start] == arr[end] {
			count++
			 end++
			start--
		}
		if count > max {
		   max = count
		   //get the start of this choise and divide by two, since
		   //we add special character
		   maxInd = (ind - maxInd) / 2
		}
	}
   return s[maxInd: maxInd+max]
}