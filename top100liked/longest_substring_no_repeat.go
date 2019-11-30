package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
/**
  Leverage the idea of sliding window. There are <start, end> pointer, initially as 0, 
  increase end pointer, when a repeat character is found, check whether the position of the character
  is  >= start and < end
	True: then start = pos, and update the longest string variable
	False: proceed 

	Always update  a character's latest position
*/
func lengthOfLongestSubstring(s string) int {
	currentLongest , start, end := 0, 0 , 0
	dict := make([]int, 256)
	for i := range dict {
		dict[i] = -1
	}
	for end < len(s) {
		if pos := dict[s[end]]; pos < end && pos >= start {
			currentLongest = max(end-start, currentLongest)
			start = pos + 1
		}
		dict[s[end]] = end
		end++
	}
	return max(end-start, currentLongest)
}
