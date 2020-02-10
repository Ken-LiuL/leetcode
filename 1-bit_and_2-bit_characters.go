package array

func isOneBitCharacter(bits []int) bool {
	res := false
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
		   i++
		} else if i == len(bits) - 1 {
			 res = true
		}
   }
	return res
}