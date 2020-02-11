package dynamic 

func isSubsequence(s string, t string) bool {
	 sp, tp := 0, 0
	 for sp < len(s) && tp < len(t) {
		 if s[sp] == t[tp] {
			 sp++
		 }
		 tp++
	 }
	 if sp == len(s) {
		 return true
	 }
	 return false
}