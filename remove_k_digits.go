package string



func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, k)
	
	for i := 0; i < len(num); i++ {
		for len(stack) != 0 && k > 0 && stack[len(stack) - 1] > num[i] {
			stack = stack[: len(stack) - 1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack) - k] 
	trailing := 0  
	for trailing < len(stack) && stack[trailing] == '0' {
	   trailing++
	}
   if len(stack[trailing:]) == 0 {
	   return "0"
   }
   return string(stack[trailing:])
}