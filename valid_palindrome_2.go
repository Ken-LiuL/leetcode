package str

func isPalindrome(s string, l , r int) bool {
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return  true
}

func validPalindrome(s string) bool {
	start, end := 0, len(s) - 1
	for start < end {
		if  s[start] == s[end] {
			start++
			end--
		} else {
			return  isPalindrome(s, start, end-1) || isPalindrome(s, start+1, end)
		}
	}
	return true

}