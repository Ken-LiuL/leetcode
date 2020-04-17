package pointer


const BACK = rune('#')

func getRealStr(s string) string {
	stack := make([]rune, 0, len(s))
	l := 0
	for _, c := range s {
	   if c == BACK {
           if l > 0 {
		    stack = stack[:l-1]
            l -= 1
           }
	   } else {
		   stack = append(stack, c)
		   l += 1
	   }
	} 
	return string(stack)
}

func backspaceCompare(S string, T string) bool {
	 return  getRealStr(S) == getRealStr(T)
}