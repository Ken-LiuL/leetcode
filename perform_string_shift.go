package  pointers


//strings as loop linked list
//to left is clockwise, to right is counter-clockwise
func stringShift(s string, shift [][]int) string {
	  start, size := 0, len(s)
      for _, s := range shift { 
		  //clockwise
		  if s[0] == 0 {
				start += s[1]
				start %=  size
		  } else {
			//counter clockwise
				start -= s[1]
				start %= size
				if start < 0 {
					start += size
				}
		  }
	  }
	  return s[start:]+s[:start]
}