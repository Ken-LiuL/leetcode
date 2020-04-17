package string


func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int)  int {
	if a < b {
		return a
	}
	return b
}

func maxRepOpt1(text string) int {
	freq := make([]int, 26)
	start, end, maxChar, maxRep, res := 0, 0, 0 , 0, 0

	for end < len(text) {
		freq[text[end] - 'a']++
		if freq[text[end] - 'a'] > maxRep {
			maxRep = freq[text[end] - 'a']
            maxChar = int(text[end] - 'a')
		}
		if end - start + 1 - maxRep > 1 {
			freq[text[start] - 'a']--
			start++
		}
		res = max(end-start+1, res)
        end++
	}
    cnt := 0
    for i := 0; i < len(text); i++ {
        if int(text[i] - 'a') == maxChar {
            cnt++
        }
    }
	return min(res, cnt)
}