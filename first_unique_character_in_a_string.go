package string


func firstUniqChar(s string) int {
	characters := make([]int, 26)
	ind := 0
	if len(s) == 0 {
		return -1
	}
	for i := 0; i < len(s); i++ {
		characters[s[i]-'a']++
		if characters[s[i]-'a'] > 1 && s[i] == s[ind] {
            ind++
			for ind < len(s) && characters[s[ind] - 'a'] >  1{
				ind++
			}
			if ind >= len(s) {
				return -1
			}
		}
	}
	return ind
}