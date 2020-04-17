package array


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	freq := make([]int, 26)
	start, end, mostFreq := 0, 0, 0
	longest := 0
	for end < len(s) {
		freq[s[end]-'A']++ 
		mostFreq = max(mostFreq, freq[s[end]-'A'])
		if  mostFreq + k < end - start + 1  {
			freq[s[start]-'A']--
			start++ 
		} else if end - start + 1 >  longest {
			longest = end - start + 1
		}   
		end++
   }
   return  longest
}