package str


//we only need to compare the max string with the min string
//since the longest common prefix would be the min string
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1  {
		return strs[0]
	}
	maxStr, minStr := strs[0], strs[0]
	for _, n := range strs[1:] {
		if n > maxStr {
			maxStr = n
		}
		if n < minStr {
			minStr = n
		}
	}
	for ind, r := range minStr {
        if rune(maxStr[ind]) != r {
			return minStr[:ind]
		}
	}
	return minStr
}