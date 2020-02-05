package str

func lengthOfLastWord(s string) int {
	inWord, startOfWord, endOfWord := false, -1, -1
	for i:=0; i < len(s)+1; i++ {
	   if i == len(s)   {
		   if inWord {
			 endOfWord = i 
		   }
		   break
	   }
	   if s[i] == ' ' {
		   if inWord {
			   endOfWord = i 
			   inWord = false
		   }
	   } else {
		   if !inWord {
			   startOfWord = i
			   inWord = true
		   } 
	   }
   }
   if endOfWord == -1 {
	   return 0
   } else {
	   return  endOfWord- startOfWord
   }
}