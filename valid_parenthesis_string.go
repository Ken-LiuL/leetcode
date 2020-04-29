package string

func checkValidString(s string) bool {
	 cmin, cmax :=0, 0
	 for   i := 0; i < len(s); i++ {
		 if s[i] == '(' {
			cmin++
			cmax++
		 } else if s[i] == ')' {
			cmin--
			cmax--
		 } else {
			cmax++
			cmin--
		 }
		 if cmax < 0 {
			 return false
		 }
		 if cmin < 0 {
			 cmin = 0
		 }
	 }
	 return cmin == 0
}