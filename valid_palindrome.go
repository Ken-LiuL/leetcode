package str

func convertCase(s byte) (byte, bool) {
	if  'a' <= s && s <= 'z' {
		return s, true
	}
	if 'A' <= s && s <= 'Z' {
		return s+32, true
	}
    if '0' <= s && s <= '9' {
        return s, true
    }
	return 0, false
}

//ignore non-alphabetic character
//use two pointer
func isPalindrome(s string) bool {
	 start, end := 0, len(s) - 1
	 for start < end {
		 startCase, startOk := convertCase(s[start])
		 endCase, endOk := convertCase(s[end])
		 if !startOk {
			 start++
		 } else if !endOk {
			 end--
		 } else if startCase == endCase {
			start++
			end--
		 } else {
			 return false
		 }
	 }
	 return true
}