package string 

func countSegments(s string) int {
	   inWord, cnt := false, 0
	   for i := 0; i < len(s); i++ {
			if s[i] == ' ' && inWord {
				 cnt++
				 inWord = false
			} else if s[i] != ' ' && !inWord {
				inWord = true
			}
	   }
	   return cnt
}