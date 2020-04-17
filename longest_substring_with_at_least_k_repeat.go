package string

import "strings"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubstring(s string, k int) int {
    if  k == 0 || k == 1 || len(s) == 0 {
        return len(s)
	}
	cache := make([]int, 26)
	badInd := -1
	for  i := 0; i < len(s); i++ {
		cache[s[i]-'a']++
	}
	for ind, v := range cache {
		if v < k && v != 0 {
			badInd = ind
			break
		}
	}
	if badInd == -1 {
        return len(s)
	}
    pivot := strings.IndexByte(s, byte(badInd+'a'))
    return max(longestSubstring(s[:pivot], k), longestSubstring(s[pivot+1:],k))
}